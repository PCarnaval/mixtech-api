package handler

import (
	"net/http"

	"github.com/PCarnaval/mixtech-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpening(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error fetching openings from database")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
