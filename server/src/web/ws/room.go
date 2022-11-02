package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/usecase"
)

type RoomWsHandler struct {
	uc   usecase.IRoomUseCase
	Hubs *Hubs
}

func NewRoomWsHandler(uc usecase.IRoomUseCase, hubs *Hubs) *RoomWsHandler {
	return &RoomWsHandler{
		uc:   uc,
		Hubs: hubs,
	}
}

// ConnectWsRoomはwebsocket通信でserver内のRoomsオブジェクトとデータをやり取りします
func (h *RoomWsHandler) ConnectWsRoom(ctx *gin.Context) {
	roomId := (entity.RoomId)(ctx.Param("room"))
	hub := h.getOrRegisterHub(roomId)
	h.uc.CheckExistsRoomAndInit(roomId)
	h.serveWsConnOfHub(hub, ctx.Writer, ctx.Request)
}

// getOrRegisterHubはHubがすでに登録されていた場合既存のHubを返し、なかった場合は新規登録します
func (h *RoomWsHandler) getOrRegisterHub(roomId entity.RoomId) *Hub {
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
