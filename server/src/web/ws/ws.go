package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/jphacks/F_2205/server/src/domain/entity"

	"github.com/gorilla/websocket"
)

// receiveEventInfoFromConnはクライアントからEvent情報が送られてきたとき、
// Eventごとに処理を行い、新たなRoom情報をBroadcastRoomInfoに書き込みます
func (h *RoomWsHandler) receiveEventInfoFromConn(c *Client) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.setConnectionConfig()

	for {
		e := entity.Event{}
		// ここで処理をまってるっぽい
		if err := c.Conn.ReadJSON(&e); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error : %v", err)
			}
			break
		}
		// eventを実行して、最新のroomオブジェクトを返す
		room,err := h.uc.ExecEventOfEventType(e,c.Hub.RoomId)
		if err != nil {
			log.Println("ExecEventOfEventType Error :", err)
		}
		// FocusMemberに最新の情報をいれる
		room.FocusMembers = h.uc.GetFocusMembersOfRoomId(c.Hub.RoomId)
		c.Hub.BroadcastRoomInfo <- room
	}
}

// sendRoomInfoToAllClientsはBroadcastRoomInfoに入ってきたRoomの情報を
// Hubに登録されたすべてのクライアントに送ります
func (h *RoomWsHandler) sendRoomInfoToAllClients(c *Client) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case room, ok := <-c.SendRoomInfo:
			c.setWriteDeadline()
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.Conn.WriteJSON(room); err != nil {
				log.Println("Error : write json failed")
			}

		case <-ticker.C:
			c.setWriteDeadline()
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWsConnOfHubはコネクションをwebsocket通信にアップグレードし、
// Clientオブジェクトを用意してサーバーを起動します
func (h *RoomWsHandler) serveWsConnOfHub(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error : ", err)
		return
	}

	client := &Client{
		Hub:          hub,
		Conn:         conn,
		SendRoomInfo: make(chan *entity.Room),
	}

	// HubにClientを登録する
	client.Hub.Register <- client

	go h.sendRoomInfoToAllClients(client)
	go h.receiveEventInfoFromConn(client)
}
