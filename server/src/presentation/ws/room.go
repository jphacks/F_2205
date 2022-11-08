package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/service"
	"github.com/jphacks/F_2205/server/src/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type RoomWsHandler struct {
	ucRoom  usecase.IRoomUsecase
	ucEvent usecase.IEventUsecase
	Hubs    *Hubs
}

// NewRoomWsHandlerはusecaseの依存を注入しRoomWsHandler構造体を返します
func NewRoomWsHandler(ucRoom usecase.IRoomUsecase, ucEvent usecase.IEventUsecase, hubs *Hubs) *RoomWsHandler {
	return &RoomWsHandler{
		ucRoom:  ucRoom,
		ucEvent: ucEvent,
		Hubs:    hubs,
	}
}

// ConnectWsRoomはwebsocket通信でserver内のRoomsオブジェクトとデータをやり取りします
func (h *RoomWsHandler) ConnectWsRoom(ctx *gin.Context) {
	roomIdString := ctx.Param("room_id")
	roomId := service.StringToRoomId(roomIdString)

	var hub *Hub
	// もしすでにroomIdのHubが存在していた場合hubに入れる
	hub, found := h.Hubs.getExistsHubOfRoomId(roomId)

	// roomIdのHubが存在していなかったら新しく登録し、Hubを起動する
	if !found {
		hub = NewHub(roomId)
		h.Hubs.setNewHubOfRoomId(hub, roomId)
		go hub.Run()
	}
	h.ucRoom.CheckExistsRoomAndInit(roomId)
	h.serveWsConnOfHub(hub, ctx.Writer, ctx.Request)
}

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
		// ここで処理をまつ
		if err := c.Conn.ReadJSON(&e); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error : %v", err)
			}
			break
		}
		// eventを実行して、最新のroomオブジェクトを返す
		room, err := h.ucEvent.SwitchExecEventByEventType(e, c.Hub.RoomId)
		if err != nil {
			log.Println("ExecEventOfEventType Error :", err)
		}
		h.ucRoom.SetRoomLatestMemberDataOfRoomId(c.Hub.RoomId,room,e)
		c.Hub.BroadcastRoom <- room
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
// Clientオブジェクトを用意してClientを受け取ったHubに登録します
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
