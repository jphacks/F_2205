package router

import (
	"fmt"
	"time"

	"github.com/jphacks/F_2205/server/src/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

	r.engine.Use(cors.New(cors.Config{
	// アクセスを許可したいアクセス元
	AllowOrigins: []string{
		"*",
	},
	// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
	AllowMethods: []string{
		"POST",
		"GET",
		"OPTIONS",
	},
	// 許可したいHTTPリクエストヘッダ
	AllowHeaders: []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	},
	// cookieなどの情報を必要とするかどうか
	AllowCredentials: true,
	// preflightリクエストの結果をキャッシュする時間
	MaxAge: 24 * time.Hour,
	}))
}
