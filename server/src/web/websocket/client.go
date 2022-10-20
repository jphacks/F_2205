// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"log"
	"net/http"
	"time"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/usecase"

	"github.com/gorilla/websocket"
)

// readPumpはクライアントから何か送られてきたとき、broadcastに書き込みます
func readPump(c *entity.Client, focusUC *usecase.FocusUseCase) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		e := entity.Event{}
		// ここで処理をまってるっぽい
		if err := c.Conn.ReadJSON(&e); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error : %v", err)
			}
			break
		}

		// Eventのタイプによって処理を切り替える
		switch e.Type {
		case entity.NewMember:
			if err := focusUC.NewMember(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		case entity.SetFocus:
			if err := focusUC.SetFocus(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		case entity.DelFocus:
			if err := focusUC.DelFocus(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		default:
			log.Println("Error : not matched type")
		}

		// 最新のHubの状態を書き込む
		e.Focus.Members = c.Hub.Focus.Members
		c.Hub.Broadcast <- &e
	}
}

// writePumpはbroadcastに入ってきたものを、hubに登録されたクライアント全員に送ります
func writePump(c *entity.Client) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case s, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.Conn.WriteJSON(s); err != nil {
				log.Println("Error : write json failed")
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWs(hub *entity.Hub, w http.ResponseWriter, r *http.Request, focusUC *usecase.FocusUseCase) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error : ", err)
		return
	}

	client := &entity.Client{Hub: hub, Conn: conn, Send: make(chan *entity.Event)}
	client.Hub.Register <- client

	go writePump(client)
	go readPump(client, focusUC)
}
