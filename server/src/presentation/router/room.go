package router

import (
	"github.com/jphacks/F_2205/server/src/domain/service"
	"github.com/jphacks/F_2205/server/src/infrastructure/persistance"
	"github.com/jphacks/F_2205/server/src/presentation/handler"
	"github.com/jphacks/F_2205/server/src/presentation/ws"
	"github.com/jphacks/F_2205/server/src/usecase"
)

// InitRoomRouterはroomAPIのハンドラーを用意します
func (r Router) InitRoomRouter() {
	hubsStore := ws.NewHubsStore()
	rooms := service.NewRooms()

	repo := persistance.NewRoomRepository(rooms)
	ucRoom := usecase.NewRoomUsecase(repo)
	ucEvent := usecase.NewEventUsecase(repo)
	h := handler.NewRoomHandler(ucRoom)
	hWs := ws.NewRoomWsHandler(ucRoom, ucEvent, hubsStore)

	// Room API
	r.Engine.GET("/room/:room_id", h.GetRoomOfRoomId)
	r.Engine.POST("/room", h.CreateRoom)
	r.Engine.DELETE("/room/:room_id", h.DeleteRoomOfRoomId)
	r.Engine.GET("/room/sum", h.GetCountSumOfRoom) //debug用

	// WebSocket関係
	r.Engine.GET("/ws/:room_id", hWs.ConnectWsRoom)
	r.Engine.GET("/ws/conn/:room_id", hWs.GetConnCountOfRoomId) //debug用
	r.Engine.DELETE("/ws/:room_id", hWs.DeleteHubOfRoomId)
}
