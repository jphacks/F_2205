package main

import (
	"log"

	"github.com/jphacks/F_2205/server/src/infrastructure/database"
	"github.com/jphacks/F_2205/server/src/web/router"
)

func main() {
	conn, err := database.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Routerの初期化
	r := router.NewRouter()
	r.SetMiddleware()

	r.Health()
	r.NewRoomRouter(conn)

	// Routerの起動
	r.Serve()
}