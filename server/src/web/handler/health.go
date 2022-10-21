package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandlerはHealthをチェックするhandlerです
func HealthHandler(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{"health": "good"},
	)
}
