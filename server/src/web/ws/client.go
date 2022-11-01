package ws

import (
	"github.com/gorilla/websocket"
	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// Clientはwebsocketとhubをつなぐ構造体です
type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan *entity.Event // Buffered channel of outbound messages.
	RoomId entity.RoomId
}
