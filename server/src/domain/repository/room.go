package repository

import (
	"github.com/jphacks/F_2205/server/src/domain/entity"
)

type IRoomRepository interface {
	CreateRoom(room *entity.Room) (*entity.Room, error)
}
