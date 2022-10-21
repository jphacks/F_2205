package router

import (
	"github.com/jphacks/F_2205/server/src/web/handler"
)

// Healthはサーバーのヘルスチェックをするrouterです
func (r Router) Health() {
	r.Engine.GET("/health", handler.HealthHandler)
}
