// hub.goではwebsocket通信で使うHubの構造体を管理しています

package ws

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	BroadcastRoom chan *entity.Room

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	RoomId entity.RoomId
}

type Hubs map[entity.RoomId]*Hub

// NewHubは新しいHubオブジェクトを生成します
func NewHub(roomId entity.RoomId) *Hub {
	return &Hub{
		BroadcastRoom:     make(chan *entity.Room),
		Register:          make(chan *Client),
		Unregister:        make(chan *Client),
		Clients:           make(map[*Client]bool),
		RoomId:            roomId,
	}
}

// NewHubsは新たにHubsオブジェクトのポインタを返します
func NewHubs() *Hubs {
	return &Hubs{}
}

// RunはHubを起動し、待ち受け状態にします
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.SendRoomInfo)
			}
		// BroadcastRoomにroom情報が来た時、各ClientのSendRoomInfoチャネルにroom情報を送る
		case room := <-h.BroadcastRoom:
			for client := range h.Clients {
				select {
				case client.SendRoomInfo <- room:
				default:
					close(client.SendRoomInfo)
					delete(h.Clients, client)
				}
			}
		}
	}
}

// setNewHubOfRoomIdはRoomIdをKeyにHubsに新しいHubを登録します
func (h *Hubs) setNewHubOfRoomId(hub *Hub, roomId entity.RoomId) {
	(*h)[roomId] = hub
}

// getExistsHubOfRoomIdはroomIdのHubが存在するか確認し、存在した場合は取得したHubを返します
func (h *Hubs) getExistsHubOfRoomId(roomId entity.RoomId) (*Hub, bool) {
	hub, ok := (*h)[roomId]
	if !ok {
		return nil, false
	}
	return hub, true
}

// CheckAndDeleteHubOfRoomIはroomIdのHubが存在するか確認し
// 存在した場合は削除し、存在しなかった場合はエラーを返します
func (h *Hubs) CheckAndDeleteHubOfRoomId(roomId entity.RoomId) error {
	if _, found := h.getExistsHubOfRoomId(roomId); !found {
		return fmt.Errorf("Hubs.CheckAndDeleteHubOfRoomId Error : roomId not found in Hubs")
	}
	delete(*h, roomId)
	return nil
}
