package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

// RoomUseCaseがIRoomUsecaseを満たしているか確認します
var _ IRoomUsecase = &RoomUseCase{}

// RoomUseCaseはRoomに関するユースケースです
type RoomUseCase struct {
	roomRepo repository.IRoomRepository
}

// IRoomUsecaseはRoomに関するユースケースのインターフェースです
type IRoomUsecase interface {
	CreateRoom(room *entity.Room) (*entity.Room, error)
}

// NewRoomUseCaseはRoomUseCaseのオブジェクトのポインタを生成する関数です
func NewRoomUseCase(r repository.IRoomRepository) *RoomUseCase {
	return &RoomUseCase{
		roomRepo: r,
	}
}

// CreateRoomはrepositoryの関数を呼び出し、新しいroomを保存します
func (u *RoomUseCase) CreateRoom(room *entity.Room) (*entity.Room, error) {
	if room.Name == "" {
		return nil, fmt.Errorf("RoomUseCase.CreateRoom Name Error : room name required")
	}
	room, err := u.roomRepo.CreateRoom(room)
	if err != nil {
		return nil, fmt.Errorf("RoomUseCase.CreateRoom CreateRoom Error : %w", err)
	}
	return room, nil
}
