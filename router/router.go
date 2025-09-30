package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	// Creates a Gin router with default configs:
	router := gin.Default()

	// Initialize routes
	InitializeRoutes(router)

	router.Run(":8080")
}
