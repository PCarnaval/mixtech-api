package main

import (
	"log"

	"github.com/PCarnaval/mixtech-api/config"
	"github.com/joho/godotenv"

	//	"github.com/PCarnaval/mixtech-api/middleware/""
	"github.com/PCarnaval/mixtech-api/router"
)

var (
	logger *config.Logger
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env não encontrado, usando variáveis do sistema")
	}

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error("config initialization error: %v", err)
		return
	}

	router.InitializeRouter()

}
