package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/PCarnaval/mixtech-api/config"
	"github.com/PCarnaval/mixtech-api/router"
)

var (
	logger *config.Logger
)

func main() {

	payload := []byte("Conteúdo da mensagem a ser assinada")
	key := []byte("MinhaChaveSecretaSuperSegura")

	signature := HMACSign(payload, key)
	fmt.Println("Assinatura:", signature)

	if HMACVerify(payload, key, signature) {
		fmt.Println("Assinatura válida!")
		return
	}

	fmt.Println("Assinatura inválida!")

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error("config initialization error: %v", err)
		return
	}

	router.InitializeRouter()

}

func HMACSign(payload, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func HMACVerify(payload, key []byte, receivedSig string) bool {
	rMac, err := base64.StdEncoding.DecodeString(receivedSig)
	if err != nil {
		return false
	}

	eMac := hmac.New(sha256.New, key)
	eMac.Write(payload)

	return hmac.Equal(eMac.Sum(nil), rMac)
}
