package router

import (
	"github.com/jphacks/F_2205/server/src/domain/service"
	"github.com/jphacks/F_2205/server/src/infrastructure/persistance"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/web/handler"
	"github.com/jphacks/F_2205/server/src/web/ws"
)

func (r Router) InitWsRoomRouter() {

	hubs := ws.NewHubs()
	rooms := service.NewRooms()

	repo := persistance.NewRoomRepository(rooms)
	uc := usecase.NewRoomUseCase(repo)
	h := handler.NewRoomHandler(uc, hubs)
	hWs := ws.NewRoomWsHandler(uc, hubs)

	r.Engine.DELETE("/ws/:room", h.DeleteHubOfRoomId)
	r.Engine.GET("/ws/:room", hWs.ConnectWsRoom)
}