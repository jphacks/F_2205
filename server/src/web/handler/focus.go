package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/usecase"
)

type FocusHandler struct {
	FocusUC usecase.IFocusUseCase
}

func NewFocusHandler(u usecase.IFocusUseCase) *FocusHandler {
	return &FocusHandler{
		FocusUC: u,
	}
}

func (h *FocusHandler) WsFocus(ctx *gin.Context) {
	roomId := (entity.RoomId)(ctx.Param("room"))
	hub := h.FocusUC.GetOrRegisterHub(roomId)
	h.serveWs(hub, ctx.Writer, ctx.Request)
}

func (h *FocusHandler) DeleteRoomHub(ctx *gin.Context) {
	roomId := (entity.RoomId)(ctx.Param("room"))
	if err := h.FocusUC.CheckHubExists(roomId); err != nil {
		ctx.JSON(
			http.StatusOK,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"ok": "delete room successful"},
	)
}

// readPumpはクライアントから何か送られてきたとき、broadcastに書き込みます
func (h *FocusHandler) readPump(c *entity.Client) {
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
			if err := h.FocusUC.NewMember(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		case entity.SetFocus:
			if err := h.FocusUC.SetFocus(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		case entity.DelFocus:
			if err := h.FocusUC.DelFocus(c.Hub.RoomId, e.Info); err != nil {
				log.Println("Error : ", err)
			}
		case entity.DelAllFocus:
			if err := h.FocusUC.DelAllFocus(c.Hub.RoomId, e.Info); err != nil {
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
func (h *FocusHandler) writePump(c *entity.Client) {
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

func (h *FocusHandler) serveWs(hub *entity.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error : ", err)
		return
	}

	client := &entity.Client{Hub: hub, Conn: conn, Send: make(chan *entity.Event)}
	client.Hub.Register <- client

	go h.writePump(client)
	go h.readPump(client)
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// upgraderはHTTP通信からwebsocket プロトコルにアップグレードします
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
