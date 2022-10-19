package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/F_2205/server/src/config"
)

type Router struct {
	engine *gin.Engine
}

// NewRouterは新しいRouterを初期化し構造体のポインタを返します
func NewRouter() *Router {
	e := gin.New()
	return &Router{
		engine: e,
	}
}

// Serveはhttpサーバーを起動します
func (r *Router) Serve() {
	r.engine.Run(fmt.Sprintf(":%s", config.Port()))
}

// NewMiddlewareはmiddlewareを用意します
func (r *Router) SetMiddleware() {
	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())
}
