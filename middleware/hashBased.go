package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	HeaderClientID  = "X-Client-Id"
	HeaderSignature = "X-Signature"
	HeaderTimestamp = "X-Timestamp"
	AllowedDriftSec = 300 // 5 minutos
)

// Middleware HMACAuth
func HMACAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := r.Header.Get(HeaderClientID)
		sigHex := r.Header.Get(HeaderSignature)
		tsStr := r.Header.Get(HeaderTimestamp)

		if clientID == "" || sigHex == "" || tsStr == "" {
			http.Error(w, "missing auth headers", http.StatusUnauthorized)
			return
		}

		tsInt, err := strconv.ParseInt(tsStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid timestamp", http.StatusUnauthorized)
			return
		}
		reqTime := time.Unix(tsInt, 0)
		if !validTimestamp(reqTime) {
			http.Error(w, "timestamp drift too large", http.StatusUnauthorized)
			return
		}

		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		canonical := buildCanonicalMessage(clientID, tsStr, r.Method, r.URL.Path, bodyBytes)

		secret, err := getSecretForClient(clientID)
		if err != nil {
			http.Error(w, "unknown client", http.StatusUnauthorized)
			return
		}

		expectedMAC := computeHMAC([]byte(canonical), []byte(secret))

		providedMAC, err := hex.DecodeString(strings.TrimSpace(sigHex))
		if err != nil {
			http.Error(w, "invalid signature format", http.StatusUnauthorized)
			return
		}

		if !hmac.Equal(expectedMAC, providedMAC) {
			http.Error(w, "signature mismatch", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validTimestamp(t time.Time) bool {
	now := time.Now().UTC()
	diff := now.Sub(t)
	if diff < 0 {
		diff = -diff
	}
	return diff <= time.Duration(AllowedDriftSec)*time.Second
}

func buildCanonicalMessage(clientID, ts, method, path string, body []byte) string {
	return strings.Join([]string{
		clientID,
		ts,
		strings.ToUpper(method),
		path,
		string(body),
	}, ":")
}

func computeHMAC(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

func getSecretForClient(clientID string) (string, error) {
	envKey := "HMAC_SECRET_" + clientID
	if v := os.Getenv(envKey); v != "" {
		return v, nil
	}
	return "", errors.New("secret not found")
}
