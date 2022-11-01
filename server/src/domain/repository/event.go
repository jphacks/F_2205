package repository

import "github.com/jphacks/F_2205/server/src/domain/entity"

type IEventRepository interface {
	NewMember(roomId entity.RoomId, newMemberName entity.Name) error
	SetFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error
	DelFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error
	DelAllFocus(roomId entity.RoomId, from entity.Name) error
	GetMemberOfRoomId(roomId entity.RoomId) entity.Members
	CheckExistsEventAndInit(roomId entity.RoomId)
	DeleteEventOfRoomId(roomId entity.RoomId)
}
