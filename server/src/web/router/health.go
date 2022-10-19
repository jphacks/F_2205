package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthはサーバーのヘルスチェックをするハンドラーです
func (r Router) Health() {
	r.engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"health": "good"})
	})
}
