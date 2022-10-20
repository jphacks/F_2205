package main

import (
	"github.com/jphacks/F_2205/server/src/web/router"
)

func main() {

	// Routerの初期化
	r := router.NewRouter()
	r.SetMiddleware()

	r.Health()
	r.Ws()

	// Routerの起動
	r.Serve()
}
