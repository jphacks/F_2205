package handler

import (
	"net/http"

	"github.com/jphacks/F_2205/server/src/domain/service"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/utils/json"

	"github.com/gin-gonic/gin"
)

// TODO handlerがHubsに依存しているのが気になる。
type RoomHandler struct {
	uc usecase.IRoomUsecase
}

// NewRoomHandlerはRoomHandler構造体のポインタを返します
func NewRoomHandler(uc usecase.IRoomUsecase) *RoomHandler {
	return &RoomHandler{
		uc: uc,
	}
}

// GetRoomOfRoomIdは指定したRoomIdにRoomを取得するAPIです
func (h *RoomHandler) GetRoomOfRoomId(ctx *gin.Context) {
	roomIdString := ctx.Param("room_id")
	roomId := service.StringToRoomId(roomIdString)

	room, err := h.uc.GetRoomOfRoomId(roomId)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	roomJson := json.RoomEntityToJson(room)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": roomJson},
	)
}

// CreateRoomは新しいRoomを作成するハンドラーです
func (h *RoomHandler) CreateRoom(ctx *gin.Context) {
	roomInfo, err := h.uc.CreateRoomNumber()
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	h.uc.CheckExistsRoomAndInit(roomInfo.Id)
	roomInfoJson := json.RoomInfoEntityToJson(roomInfo)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": roomInfoJson},
	)
}

// DeleteRoomOfRoomIdは受け取ったroom_idのRoomを削除するハンドラーです
func (h *RoomHandler) DeleteRoomOfRoomId(ctx *gin.Context) {
	roomIdString := ctx.Param("room_id")
	roomId := service.StringToRoomId(roomIdString)

	// Roomの削除
	// TODO 見つからなかった時にエラーを返す処理にしたい
	h.uc.DeleteRoomOfRoomId(roomId)

	ctx.JSON(
		http.StatusOK,
		gin.H{"ok": "delete room of roomId successful"},
	)
}

// GetSumOfRoomは存在する部屋の数を返すハンドラーです
func (h *RoomHandler) GetCountSumOfRoom(ctx *gin.Context) {
	cnt := h.uc.GetSumOfRoom()
	cntRoomJson := json.CreateCountRoomJson(cnt)

	ctx.JSON(
		http.StatusOK,
		gin.H{"data": cntRoomJson},
	)
}
