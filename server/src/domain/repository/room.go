package repository

import "github.com/jphacks/F_2205/server/src/domain/entity"

type IRoomRepository interface {
	// room.go
	AddNewMemberOfRoomId(roomId entity.RoomId, member *entity.Member, peerId entity.PeerId) error
	CheckExistsRoomAndInit(roomId entity.RoomId)
	DeleteRoomOfRoomId(roomId entity.RoomId)
	GetExistsRoomOfRoomId(roomId entity.RoomId) (*entity.Room, bool)
	GetSumOfRoom() int
	GetMembersOfRoomId(roomId entity.RoomId) entity.Members

	// room_focus.go
	AddNewFocusMemberOfRoomId(roomId entity.RoomId, newPeerId entity.PeerId) error
	SetMemberFocusOfRoomId(roomId entity.RoomId, from entity.PeerId, to entity.PeerId) error
	DelMemberFocusOfRoomId(roomId entity.RoomId, from entity.PeerId, to entity.PeerId) error
	DelAllMemberFocusOfRoomId(roomId entity.RoomId, from entity.PeerId) error
	GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers

	// room_restroom.go
	SetMemberRestRoomStateOfRoomId(roomId entity.RoomId, peerId entity.PeerId, isRestRoom bool) error
}
