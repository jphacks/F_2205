package repository

import "github.com/jphacks/F_2205/server/src/domain/entity"

type IRoomRepository interface {
	// room.go
	AddNewMemberOfRoomId(roomId entity.RoomId, member *entity.Member) error 
	CheckExistsRoomAndInit(roomId entity.RoomId)
	DeleteRoomOfRoomId(roomId entity.RoomId)
	GetExistsRoomOfRoomId(roomId entity.RoomId) (*entity.Room, bool)
	GetSumOfRoom() int
	GetMembersOfRoomId(roomId entity.RoomId) entity.Members

	// room_focus.go
	AddNewFocusMemberOfRoomId(roomId entity.RoomId, newMemberName entity.Name) error
	SetMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name, to entity.Name) error
	DelMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name, to entity.Name) error
	DelAllMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name) error
	GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers
}
