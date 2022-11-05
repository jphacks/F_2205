package router

import (
	"github.com/jphacks/F_2205/server/src/domain/service"
	"github.com/jphacks/F_2205/server/src/infrastructure/persistance"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/web/handler"
	"github.com/jphacks/F_2205/server/src/web/ws"
)

// InitRoomRouterはroomAPIのハンドラーを用意します
func (r Router) InitRoomRouter() {
	hubs := ws.NewHubs()
	rooms := service.NewRooms()

	repo := persistance.NewRoomRepository(rooms)
	ucRoom := usecase.NewRoomUsecase(repo)
	ucEvent := usecase.NewEventUsecase(repo)
	h := handler.NewRoomHandler(ucRoom, hubs)
	hWs := ws.NewRoomWsHandler(ucRoom, ucEvent, hubs)

	// Room API
	r.Engine.GET("/room/sum", h.GetCountSumOfRoom)
	r.Engine.POST("/room", h.CreateRoom)
	r.Engine.DELETE("/room/:room_id", h.DeleteHubAndRoomOfRoomId)

	// WebSocket通信
	r.Engine.GET("/ws/:room_id", hWs.ConnectWsRoom)
}
