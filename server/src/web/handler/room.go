package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/usecase"
)

// RoomHandlerはhttpをルーティングするハンドラーに関する構造体です
type RoomHandler struct {
	roomUC usecase.IRoomUsecase
}

// NewRoomHandlerはRoomHandlerのオブジェクトのポインタを生成する関数です
func NewRoomHandler(u usecase.IRoomUsecase) *RoomHandler {
	return &RoomHandler{
		roomUC: u,
	}
}

// CreateRoomは新しいタスクを保存するハンドラーです
func (h *RoomHandler) CreateRoom(ctx *gin.Context) {
	var b roomJson
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Errorf("RoomHandler.CreateRoom ShouldBindJSON Error : %w", err).Error(),
			},
		)
		return
	}
	room := roomJsonToEntity(&b)

	room, err := h.roomUC.CreateRoom(room)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Errorf("RoomHandler.CreateRoom CreateRoom Error : %w", err).Error(),
			},
		)
		return
	}

	roomResp := roomEntityToJson(room)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data": roomResp,
		},
	)
}

// roomJsonはroomの情報をJSONにバインドするための構造体です
type roomJson struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
}

// roomEntityToJsonはentity.RoomをroomJson型に変換します
func roomEntityToJson(room *entity.Room) *roomJson {
	return &roomJson{
		Id:     room.Id,
		Name:   room.Name,
	}
}

// roomEntityListToJsonはentity.RoomのリストをroomJsonのリストに変換します
func roomEntityListToJson(rooms entity.Rooms) []*roomJson {
	var roomJsonList []*roomJson
	for _, t := range rooms {
		roomJsonList = append(roomJsonList, roomEntityToJson(t))
	}
	return roomJsonList
}

// roomJsonToEntityはroomJson型のオブジェクトをentity.Roomに変換します
func roomJsonToEntity(roomJson *roomJson) *entity.Room {
	return &entity.Room{
		Id:     roomJson.Id,
		Name:   roomJson.Name,
	}
}
