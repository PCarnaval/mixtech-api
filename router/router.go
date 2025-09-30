package router

import "github.com/gin-gonic/gin"

func InitializeRoutes() {
	// Creates a Gin router with default configs:
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
