package service

import "github.com/jphacks/F_2205/server/src/domain/entity"

// NewRoomsは新たにRoomsオブジェクトを生成して返します
func NewRooms() *entity.Rooms {
	return &entity.Rooms{}
}

// StringToRoomIdはstring型の文字列をRoomId型に変換します
func StringToRoomId(roomIdString string) entity.RoomId {
	return (entity.RoomId(roomIdString))
}