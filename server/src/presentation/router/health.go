package router

import (
	"github.com/jphacks/F_2205/server/src/presentation/handler"
)

// Healthはサーバーのヘルスチェックをするrouterです
func (r Router) InitHealthRouter() {
	r.Engine.GET("/health", handler.HealthHandler)
}
