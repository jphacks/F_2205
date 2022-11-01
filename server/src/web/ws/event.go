package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type EventWsHandler struct {
	uc   usecase.IEventUseCase
	Hubs *Hubs
}

func NewEventWsHandler(uc usecase.IEventUseCase, hubs *Hubs) *EventWsHandler {
	return &EventWsHandler{
		uc:   uc,
		Hubs: hubs,
	}
}

// ConnectWsEventはwebsocket通信でserver内のEventsオブジェクトとデータをやり取りします
func (h *EventWsHandler) ConnectWsEvent(ctx *gin.Context) {
	roomId := (entity.RoomId)(ctx.Param("room"))
	hub := h.getOrRegisterHub(roomId)
	h.uc.CheckExistsEventAndInit(roomId)
	h.serveWs(hub, ctx.Writer, ctx.Request)
}

// readPumpはクライアントから何か送られてきたとき、broadcastに書き込みます
func (h *EventWsHandler) readPump(c *Client) {
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
			if err := h.uc.NewMember(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		case entity.SetFocus:
			if err := h.uc.SetFocus(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		case entity.DelFocus:
			if err := h.uc.DelFocus(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		case entity.DelAllFocus:
			if err := h.uc.DelAllFocus(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		default:
			log.Println("Error : not matched type")
		}

		// 指定したroomIdのMember情報をeventに入れる
		e.Focus.Members = h.uc.GetMemberOfRoomId(c.Hub.RoomId)
		c.Hub.Broadcast <- &e
	}
}

// writePumpはbroadcastに入ってきたものを、hubに登録されたクライアント全員に送ります
func (h *EventWsHandler) writePump(c *Client) {
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

// serveWsはwebsocketのサーバーを起動します
func (h *EventWsHandler) serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error : ", err)
		return
	}

	client := &Client{
		Hub:  hub,
		Conn: conn,
		Send: make(chan *entity.Event),
	}
	client.Hub.Register <- client

	go h.writePump(client)
	go h.readPump(client)
}

// getOrRegisterHubはHubがすでに登録されていた場合既存のHubを返し、なかった場合は新規登録します
func (h *EventWsHandler) getOrRegisterHub(roomId entity.RoomId) *Hub {
	// もしすでにroomIdのHubが存在していた場合hubに代入する
	hub, found := h.Hubs.getExistsHubOfRoomId(roomId)

	// roomIdのHubが存在していなかったら新しく登録する
	if !found {
		hub = NewHub(roomId)
		h.Hubs.setNewHubOfRoomId(hub, roomId)
		go hub.Run()
	}
	return hub
}
