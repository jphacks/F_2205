package main

import (
	"github.com/jphacks/F_2205/server/src/web/router"
)

func main() {

	// Routerの初期化
	r := router.NewRouter()

	// Routerの登録
	r.Health()
	r.Ws()

	// Routerの起動
	r.Serve()
}
