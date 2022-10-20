package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/infrastructure/hub"
	"github.com/jphacks/F_2205/server/src/infrastructure/persistance"
	"github.com/jphacks/F_2205/server/src/usecase"
	"github.com/jphacks/F_2205/server/src/web/websocket"
)

// Healthはサーバーのヘルスチェックをするハンドラーです
func (r Router) Ws() {

	// TODO inithub
	hubs := &entity.Hubs{}
	focusRepo := persistance.NewFocusRepository(hubs)
	focusUC := usecase.NewFocusUseCase(focusRepo)

	// TODO handler内でhubsを使っているのおかしい
	r.engine.GET("/ws/:room", func(ctx *gin.Context) {
		roomId := (entity.RoomId)(ctx.Param("room"))

		var h *entity.Hub
		// hubsに登録されているか確認
		// TODO checkする専用の関数を用意する必要ありそう
		if found, ok := (*hubs)[roomId]; ok {
			// 登録されていたら既存のものを利用
			h = found
		} else {
			// 登録されていなかったら新しく用意する
			h = hub.NewHub(roomId)
			(*hubs)[roomId] = h
			go h.Run()
		}
		websocket.ServeWs(h, ctx.Writer, ctx.Request, focusUC)
	})

	r.engine.DELETE("/ws/:room", func(ctx *gin.Context) {
		roomId := (entity.RoomId)(ctx.Param("room"))

		if _, ok := (*hubs)[roomId]; !ok {
			// 登録されていなかったら
			ctx.JSON(http.StatusOK,gin.H{"error":"room not found"})
			return
		}
		// 登録されていたら部屋を削除
		delete(*hubs,roomId)
		ctx.JSON(http.StatusOK,gin.H{"ok":"delete room successful"})
	})
}
