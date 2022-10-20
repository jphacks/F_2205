package repository

import "github.com/jphacks/F_2205/server/src/domain/entity"

type IFocusRepository interface {
	NewMember(roomId entity.RoomId, newMemberName entity.Name) error
	SetFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error
}
