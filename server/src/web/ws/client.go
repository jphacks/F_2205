// client.goではwebsocket通信で使うClientの構造体を管理しています

package ws

import (
	"time"

	"github.com/gorilla/websocket"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// Clientはwebsocketとhubをつなぐ構造体です
type Client struct {
	Hub          *Hub
	Conn         *websocket.Conn
	SendRoomInfo chan *entity.Room // Buffered channel of outbound messages.
}

// setConnectionConfigはClientのコネクション情報に関する設定をします
func (c *Client) setConnectionConfig() {
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
}

// setWriteDeadlineはClientの書き込み時間のデッドラインを設定します
func (c *Client) setWriteDeadline() {
	c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
}
