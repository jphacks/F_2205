package handler

import (
	"net/http"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/web/ws"

	"github.com/gin-gonic/gin"
)

// TODO handlerがHubsに依存しているのが気になる。
type RoomHandler struct {
	uc   usecase.IRoomUseCase
	Hubs *ws.Hubs
}

func NewRoomHandler(uc usecase.IRoomUseCase, hubs *ws.Hubs) *RoomHandler {
	return &RoomHandler{
		uc:   uc,
		Hubs: hubs,
	}
}

// GetSumOfRoomは存在する部屋の数を返すハンドラーです
func (h *RoomHandler) GetCountSumOfRoom(ctx *gin.Context) {
	cnt := h.uc.GetSumOfRoom()
	cntJson := createCountRoomJson(cnt)

	ctx.JSON(
		http.StatusOK,
		gin.H{"data": cntJson},
	)
}

func (h *RoomHandler) CreateRoom(ctx *gin.Context) {
	roomInfo, err := h.uc.CreateRoom()
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	h.uc.CheckExistsRoomAndInit(roomInfo.Id)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": roomInfo},
	)
}

func (h *RoomHandler) DeleteHubAndRoomOfRoomId(ctx *gin.Context) {
	roomId := (entity.RoomId)(ctx.Param("room_id"))

	// Roomの削除
	h.uc.DeleteRoomOfRoomId(roomId)

	// Hubの削除
	// TODO Roomを正常に削除できても、Hubを作ってなかったらエラーが出るので、他の処理方法を考えたい
	if err := h.Hubs.CheckAndDeleteHubOfRoomId(roomId); err != nil {
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

type countRoomJson struct {
	Count int `json:"count"`
}

func createCountRoomJson(count int) countRoomJson {
	return countRoomJson{
		Count: count,
	}
}
