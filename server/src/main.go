package main

import (
	"github.com/jphacks/F_2205/server/src/presentation/router"
)

func main() {

	// Routerの初期化
	r := router.NewRouter()

	// Routerの登録
	r.InitHealthRouter()
	r.InitRoomRouter()

	// Routerの起動
	r.Serve()
}
