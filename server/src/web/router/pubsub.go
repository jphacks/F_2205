package router

import (
	"github.com/jphacks/F_2205/server/src/web/pubsub"
)

func (r Router) NewPubsubRouter() {
	g := r.engine.Group("ws")
	g.GET("/room/:roomId", pubsub.PubSubHandler)
}

