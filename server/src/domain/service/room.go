package service

import (
	"sync"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// NewRoomsは新たにRoomsオブジェクトを生成して返します
func NewRooms() *entity.Rooms {
	return &entity.Rooms{}
}

// NewRoomは新たにRoomsオブジェクトを生成して返します
func NewRoom() *entity.Room {
	return &entity.Room{
		MembersStore: &entity.MembersStore{
			Members: entity.Members{},
			Mu:      sync.RWMutex{},
		},
		FocusMembersStore: &entity.FocusMembersStore{
			FocusMembers: entity.FocusMembers{},
			Mu:           sync.RWMutex{},
		},
	}
}

// StringToRoomIdはstring型の文字列をRoomId型に変換します
func StringToRoomId(roomIdString string) entity.RoomId {
	return entity.RoomId(roomIdString)
}
