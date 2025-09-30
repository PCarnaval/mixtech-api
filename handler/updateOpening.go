package handler

import (
	"net/http"

	"github.com/PCarnaval/mixtech-api/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpening(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Error("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Name != "" {
		opening.Name = request.Name
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Email != "" {
		opening.Email = request.Email
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if err := db.Save(&opening).Error; err != nil {
		logger.Error("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening on database")
		return
	}
	sendSuccess(ctx, "update-opening", opening)
}
