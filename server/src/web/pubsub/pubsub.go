package pubsub

import (
	"context"
	"log"
	"net/http"

	"github.com/jphacks/F_2205/server/src/infrastructure/redis"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Add this lines
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 空のコンテキストを生成
var ctx = context.Background()

func PubSubHandler(ctx *gin.Context){
	roomId := ctx.Query("room")
	pubsub(ctx.Writer, ctx.Request, roomId)
}

// should handle more errors
func pubsub(w http.ResponseWriter, r *http.Request, roomId string) {
	// WebSocket通信に更新
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket connection err:", err)
		return
	}
	defer conn.Close()

	go func() {
	loop:
		for {
			pubsub := redis.Rs.Subscribe(ctx, roomId)
			log.Println("redis subscribe", roomId)
			ch := pubsub.Channel()

			// should break outer for loop if err
			for msg := range ch {
				// msg.payloadにはjoinNewMemberが入っている
				// これでfrontでデータを受け取ったことを確認し、userを取得するリクエストが走る
				err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
				if err != nil {
					log.Println("websocket write err:", err)
					break loop
				} else {
					log.Println("websocket write ok!!")
				}
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("websocket read err:", err)
			break
		}
		log.Println(string(msg))

		if err := redis.Rs.Publish(ctx, roomId, msg).Err(); err != nil {
			log.Println("redis publish err:", err)
			break
		}
		log.Println("redis publish", roomId)
	}
}
