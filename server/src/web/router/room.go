package router

import (
	"github.com/jphacks/F_2205/server/src/infrastructure/database"
	"github.com/jphacks/F_2205/server/src/infrastructure/persistance"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/web/handler"
)

// NewRoomRouterはRouterへのハンドラ登録やミドルウェアの設定をします
func (r Router) NewRoomRouter(conn *database.Conn) {
	roomRepository := persistance.NewRoomRepository(conn)
	roomUseCase := usecase.NewRoomUseCase(roomRepository)
	roomHandler := handler.NewRoomHandler(roomUseCase)

	g := r.engine.Group("/room")
	g.POST("/create", roomHandler.CreateRoom)
}
