package handler

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, status int, msg string) {
	ctx.JSON(status, gin.H{
		"message":   msg,
		"errorCode": status,
	})
}
