package router

import (

	"github.com/gin-gonic/gin"
	"github.com/jphacks/F_2205/server/src/web/websocket"
)

// Healthはサーバーのヘルスチェックをするハンドラーです
func (r Router) Ws(hub *websocket.Hub) {
	r.engine.GET("/ws", func(ctx *gin.Context) {
		websocket.ServeWs(hub, ctx.Writer,ctx.Request)
	})
}
