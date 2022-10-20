package main

import (
	// "log"

	// "github.com/jphacks/F_2205/server/src/infrastructure/database"
	"github.com/jphacks/F_2205/server/src/web/websocket"
	"github.com/jphacks/F_2205/server/src/web/router"
)

func main() {
	// dbConn, err := database.NewConn()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer func() {
	// 	err := dbConn.DB.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// Routerの初期化
	r := router.NewRouter()
	r.SetMiddleware()

	r.Health()
	// r.NewRoomRouter(dbConn)
	// r.NewPubsubRouter()

	hub := websocket.NewHub()
	go hub.Run()
	r.Ws(hub)

	// Routerの起動
	r.Serve()
}
