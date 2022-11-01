package handler

import (
	"net/http"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/web/ws"

	"github.com/gin-gonic/gin"
)

// TODO handlerがHubsに依存しているのが気になる。
type EventHandler struct {
	uc   usecase.IEventUseCase
	Hubs *ws.Hubs
}

func NewEventHandler(uc usecase.IEventUseCase, hubs *ws.Hubs) *EventHandler {
	return &EventHandler{
		uc:   uc,
		Hubs: hubs,
	}
}

func (h *EventHandler) DeleteHubOfRoomId(ctx *gin.Context) {
	roomId := (entity.RoomId)(ctx.Param("room"))
	if err := h.Hubs.CheckAndDeleteHubOfRoomId(roomId); err != nil {
		ctx.JSON(
			http.StatusOK,
			gin.H{"error": err.Error()},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{"ok": "delete hub of roomId successful"},
	)
}
