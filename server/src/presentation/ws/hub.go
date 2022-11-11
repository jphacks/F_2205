// hub.goではwebsocket通信で使うHubの構造体を管理しています

package ws

import (
	"fmt"
	"sync"

	"github.com/jphacks/F_2205/server/src/utils/json"
)

type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	BroadcastRoom chan *json.RoomJson

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	RoomId json.RoomIdJson

	// スレッドセーフ
	Mu sync.RWMutex
}

type Hubs map[json.RoomIdJson]*Hub

type HubsStore struct {
	Hubs *Hubs
	Mu   sync.RWMutex
}

// NewHubは新しいHubオブジェクトを生成します
func NewHub(roomId json.RoomIdJson) *Hub {
	return &Hub{
		BroadcastRoom: make(chan *json.RoomJson),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		Clients:       make(map[*Client]bool),
		RoomId:        roomId,
		Mu:            sync.RWMutex{},
	}
}

// NewHubsは新たにHubsオブジェクトのポインタを返します
func NewHubs() *Hubs {
	return &Hubs{}
}

func NewHubsStore() *HubsStore {
	return &HubsStore{
		Hubs: NewHubs(),
		Mu:   sync.RWMutex{},
	}
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

// SetNewHubOfRoomIdはRoomIdをKeyにHubsに新しいHubを登録します
func (h *HubsStore) SetNewHubOfRoomId(hub *Hub, roomId json.RoomIdJson) {
	h.Mu.Lock()
	defer h.Mu.Unlock()

	(*h.Hubs)[roomId] = hub
}

// GetExistsHubOfRoomIdはroomIdのHubが存在するか確認し、存在した場合は取得したHubを返します
func (h *HubsStore) GetExistsHubOfRoomId(roomId json.RoomIdJson) (*Hub, bool) {
	h.Mu.RLock()
	defer h.Mu.RUnlock()

	hub, ok := (*h.Hubs)[roomId]
	if !ok {
		return nil, false
	}
	return hub, true
}

// CheckAndDeleteHubOfRoomIはroomIdのHubが存在するか確認し
// 存在した場合は削除し、存在しなかった場合はエラーを返します
func (h *HubsStore) CheckAndDeleteHubOfRoomId(roomId json.RoomIdJson) error {
	_, found := h.GetExistsHubOfRoomId(roomId)
	if !found {
		return fmt.Errorf("Hubs.CheckAndDeleteHubOfRoomId Error : roomId not found in Hubs")
	}
	h.Mu.Lock()
	defer h.Mu.Unlock()

	delete(*h.Hubs, roomId)
	return nil
}

// GetConnCountOfRoomIdは受け取ったroomIdのRoomの
// Hubに接続しているClientsの数を返します
func (h *HubsStore) GetConnCountOfRoomId(roomId json.RoomIdJson) (int, error) {
	_, found := h.GetExistsHubOfRoomId(roomId)
	if !found {
		return 0, fmt.Errorf("hub not found (roomId:%s)", (string)(roomId))
	}
	h.Mu.RLock()
	defer h.Mu.RUnlock()

	cntConn := len((*h.Hubs)[roomId].Clients)
	return cntConn, nil
}
