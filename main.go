package main

import (
	"github.com/PCarnaval/mixtech-api/config"
	"github.com/PCarnaval/mixtech-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error("config initialization error: %v", err)
		return
	}

	router.InitializeRouter()

}
