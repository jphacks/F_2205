package router

import (
	"github.com/jphacks/F_2205/server/src/domain/service"
	"github.com/jphacks/F_2205/server/src/infrastructure/persistance"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/web/handler"
	"github.com/jphacks/F_2205/server/src/web/ws"
)

// WsFocusは
func (r Router) InitWsEventRouter() {

	hubs := ws.NewHubs()
	events := service.NewEvents()

	repo := persistance.NewEventRepository(events)
	uc := usecase.NewEventUseCase(repo)
	h := handler.NewEventHandler(uc, hubs)
	hWs := ws.NewEventWsHandler(uc, hubs)

	r.Engine.DELETE("/ws/:room", h.DeleteHubOfRoomId)
	r.Engine.GET("/ws/:room", hWs.ConnectWsEvent)
}
