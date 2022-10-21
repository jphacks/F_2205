package repository

import "github.com/jphacks/F_2205/server/src/domain/entity"

type IFocusRepository interface {
	NewMember(roomId entity.RoomId, newMemberName entity.Name) error
	SetFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error
	DelFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error
	DelAllFocus(roomId entity.RoomId, from entity.Name) error 
	GetOrRegisterHub(roomId entity.RoomId) *entity.Hub
	CheckHubExists(roomId entity.RoomId) error
}
