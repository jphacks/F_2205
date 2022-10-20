package hub

import (
	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func NewHub(roomId entity.RoomId) *entity.Hub {
	return &entity.Hub{
		Broadcast:  make(chan *entity.Event),
		Register:   make(chan *entity.Client),
		Unregister: make(chan *entity.Client),
		Clients:    make(map[*entity.Client]bool),
		Focus:      &entity.Focus{},
		RoomId: roomId,
	}
}
