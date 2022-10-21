package router

import (
	"github.com/jphacks/F_2205/server/src/infrastructure/hub"
	"github.com/jphacks/F_2205/server/src/infrastructure/persistance"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/web/handler"
)

// Healthはサーバーのヘルスチェックをするハンドラーです
func (r Router) WsFocus() {

	hubs := hub.NewHubs()
	focusRepo := persistance.NewFocusRepository(hubs)
	focusUC := usecase.NewFocusUseCase(focusRepo)
	focusHandler := handler.NewFocusHandler(focusUC)

	r.Engine.GET("/ws/:room", focusHandler.WsFocus)
	r.Engine.DELETE("/ws/:room", focusHandler.DeleteRoomHub)
}
