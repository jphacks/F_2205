package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/jphacks/F_2205/server/src/domain/service"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/utils/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type RoomWsHandler struct {
	ucRoom    usecase.IRoomUsecase
	ucEvent   usecase.IEventUsecase
	HubsStore *HubsStore
}

// NewRoomWsHandlerはusecaseの依存を注入しRoomWsHandler構造体を返します
func NewRoomWsHandler(ucRoom usecase.IRoomUsecase, ucEvent usecase.IEventUsecase, hubsStore *HubsStore) *RoomWsHandler {
	return &RoomWsHandler{
		ucRoom:    ucRoom,
		ucEvent:   ucEvent,
		HubsStore: hubsStore,
	}
}

// ConnectWsRoomはwebsocket通信でserver内のRoomsオブジェクトとデータをやり取りします
func (h *RoomWsHandler) ConnectWsRoom(ctx *gin.Context) {
	roomIdString := ctx.Param("room_id")
	roomId := service.StringToRoomId(roomIdString)
	roomIdJson := json.RoomIdEntityToJson(roomId)

	var hub *Hub
	// もしすでにroomIdのHubが存在していた場合hubに入れる
	hub, found := h.HubsStore.GetExistsHubOfRoomId(roomIdJson)

	// roomIdのHubが存在していなかったら新しく登録し、Hubを起動する
	if !found {
		hub = NewHub(roomIdJson)
		h.HubsStore.SetNewHubOfRoomId(hub, roomIdJson)
		go hub.Run()
	}
	h.ucRoom.CheckExistsRoomAndInit(roomId)
	h.serveWsConnOfHub(hub, ctx.Writer, ctx.Request, roomIdJson)
}

func (h *RoomWsHandler) DeleteHubOfRoomId(ctx *gin.Context) {
	roomIdString := ctx.Param("room_id")
	roomId := service.StringToRoomId(roomIdString)
	roomIdJson := json.RoomIdEntityToJson(roomId)

	// TODO 削除時に既存のクライアントのコネクションが残っているので直す
	// Hubの削除
	if err := h.HubsStore.CheckAndDeleteHubOfRoomId(roomIdJson); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"ok": "delete hub of roomId successful"},
	)
}

func (h *RoomWsHandler) GetConnCountOfRoomId(ctx *gin.Context) {
	roomIdString := ctx.Param("room_id")
	roomId := service.StringToRoomId(roomIdString)
	roomIdJson := json.RoomIdEntityToJson(roomId)

	cnt, err := h.HubsStore.GetConnCountOfRoomId(roomIdJson)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	cntRoomConnJson := json.CreatecountRoomConnJson(cnt)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": cntRoomConnJson},
	)
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
		eJson := json.EventJson{}
		// ここで処理をまつ
		if err := c.Conn.ReadJSON(&eJson); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error : %v", err)
			}
			break
		}

		e := json.EventJsonToEntity(eJson)
		roomId := json.RoomIdJsonToEntity(c.RoomId)

		// eventを実行して、最新のroomオブジェクトを返す
		room, err := h.ucEvent.SwitchExecEventByEventType(e, roomId)
		if err != nil {
			log.Println("ExecEventOfEventType Error :", err)
		}
		h.ucRoom.SetRoomLatestMemberDataOfRoomId(roomId, room, e)

		roomJson := json.RoomEntityToJson(room)
		c.Hub.BroadcastRoom <- roomJson
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
func (h *RoomWsHandler) serveWsConnOfHub(hub *Hub, w http.ResponseWriter, r *http.Request, roomId json.RoomIdJson) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error : ", err)
		return
	}

	client := &Client{
		Hub:          hub,
		Conn:         conn,
		SendRoomInfo: make(chan *json.RoomJson),
		RoomId:       roomId,
	}

	// HubにClientを登録する
	client.Hub.Register <- client

	go h.sendRoomInfoToAllClients(client)
	go h.receiveEventInfoFromConn(client)
}
