package entity

import "sync"

// RoomIdはRoomに割り当てられるプライマリーキーです
type RoomId string

// Roomは部屋で起こったイベントの情報やステートを保持します
type Room struct {
	EventType         EventType
	MembersStore      *MembersStore
	FocusMembersStore *FocusMembersStore
	EffectMember      EffectMember
	Sound             Sound
}

type Rooms map[RoomId]*Room

type RoomsStore struct {
	Rooms *Rooms
	Mu    sync.RWMutex
}
