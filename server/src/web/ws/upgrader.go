package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// upgraderはHTTP通信からwebsocket プロトコルにアップグレードします
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
