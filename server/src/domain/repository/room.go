package repository

import "github.com/jphacks/F_2205/server/src/domain/entity"

type IRoomRepository interface {
	AddNewMemberOfRoomId(roomId entity.RoomId, newMemberName entity.Name) error
	SetMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name, to entity.Name) error
	DelMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name, to entity.Name) error
	DelAllMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name) error
	GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers
	CheckExistsRoomAndInit(roomId entity.RoomId)
	DeleteRoomOfRoomId(roomId entity.RoomId)
	GetExistsRoomOfRoomId(roomId entity.RoomId) (*entity.Room, bool)
	GetSumOfRoom() int
}
