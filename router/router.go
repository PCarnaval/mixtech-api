package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PCarnaval/mixtech-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	// Creates a Gin router with default configs:
	router := gin.Default()

	mux := http.NewServeMux()

	// Rota de teste protegida por HMAC
	mux.Handle("/api/resource", middleware.HMACAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := r.Header.Get("X-Client-Id")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message":"Autenticação HMAC OK","client":"%s"}`, clientID)
	})))

	InitializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
