package entity

import "github.com/gorilla/websocket"

// Clientはwebsocketとhubをつなぐ構造体です
// TODO websocketに依存しているがいいのか？
type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan *Event // Buffered channel of outbound messages.
	RoomId RoomId
}
