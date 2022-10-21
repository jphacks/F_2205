package service

import "github.com/jphacks/F_2205/server/src/domain/entity"

func NewHub(roomId entity.RoomId) *entity.Hub {
	return &entity.Hub{
		Broadcast:  make(chan *entity.Event),
		Register:   make(chan *entity.Client),
		Unregister: make(chan *entity.Client),
		Clients:    make(map[*entity.Client]bool),
		Focus:      &entity.Focus{},
		RoomId:     roomId,
	}
}

func NewHubs() *entity.Hubs {
	return &entity.Hubs{}
}

func Run(h *entity.Hub) {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
