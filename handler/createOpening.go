package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpening(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Error("validation error: %v", err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Error("error creating opening: %v", err.Error())
		return
	}
}
