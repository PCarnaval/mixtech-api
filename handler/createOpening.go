package handler

import (
	"net/http"

	"github.com/PCarnaval/mixtech-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpening(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Error("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:    request.Role,
		Company: request.Company,
		Name:    request.Name,
		Email:   request.Email,
		Salary:  request.Salary,
		Remote:  *request.Remote,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Error("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
