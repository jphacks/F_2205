package handler

import (
	"net/http"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/service"
	"github.com/jphacks/F_2205/server/src/presentation/ws"
	"github.com/jphacks/F_2205/server/src/usecase"

	"github.com/gin-gonic/gin"
)

// TODO handlerがHubsに依存しているのが気になる。
type RoomHandler struct {
	uc   usecase.IRoomUsecase
}

// NewRoomHandlerはRoomHandler構造体のポインタを返します
func NewRoomHandler(uc usecase.IRoomUsecase, hubs *ws.Hubs) *RoomHandler {
	return &RoomHandler{
		uc:   uc,
	}
}

// GetSumOfRoomは存在する部屋の数を返すハンドラーです
func (h *RoomHandler) GetCountSumOfRoom(ctx *gin.Context) {
	cnt := h.uc.GetSumOfRoom()
	cntRoomJson := createCountRoomJson(cnt)

	ctx.JSON(
		http.StatusOK,
		gin.H{"data": cntRoomJson},
	)
}

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
	roomInfoJson := roomInfoEntityToJson(roomInfo)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": roomInfoJson},
	)
}

func (h *RoomHandler) DeleteRoomOfRoomId(ctx *gin.Context) {
	roomIdString := ctx.Param("room_id")
	roomId := service.StringToRoomId(roomIdString)

	// Roomの削除
	h.uc.DeleteRoomOfRoomId(roomId)
	// TODO 見つからなかった時にエラーを返す処理にしたい

	ctx.JSON(
		http.StatusOK,
		gin.H{"ok": "delete room of roomId successful"},
	)
}

type roomInfoJson struct {
	Id entity.RoomId `json:"id"`
}

// roomInfoEntityToJsonはRoomInfoエンティティをresponse用のオブジェクトに変換します
func roomInfoEntityToJson(info *entity.RoomInfo) roomInfoJson {
	return roomInfoJson{
		Id: info.Id,
	}
}

type countRoomJson struct {
	Count int `json:"count"`
}

// createCountRoomJsonはcountを受け取り、response用のオブジェクトを生成します
func createCountRoomJson(count int) countRoomJson {
	return countRoomJson{
		Count: count,
	}
}
