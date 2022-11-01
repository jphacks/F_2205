package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

var _ IEventUseCase = &EventUseCase{}

type EventUseCase struct {
	repo repository.IEventRepository
}

type IEventUseCase interface {
	NewMember(roomId entity.RoomId, info entity.Info) error
	SetFocus(roomId entity.RoomId, info entity.Info) error
	DelFocus(roomId entity.RoomId, info entity.Info) error
	DelAllFocus(roomId entity.RoomId, info entity.Info) error
	GetMemberOfRoomId(roomId entity.RoomId) entity.Members
	CheckExistsEventAndInit(roomId entity.RoomId)
	DeleteEventOfRoomId(roomId entity.RoomId)
}

func NewEventUseCase(repo repository.IEventRepository) *EventUseCase {
	return &EventUseCase{
		repo: repo,
	}
}

func (u *EventUseCase) NewMember(roomId entity.RoomId, info entity.Info) error {
	newMemberName := info.From
	if newMemberName == "" {
		return fmt.Errorf("EventUseCase.NewMember Error : from is required")
	}
	return u.repo.NewMember(roomId, newMemberName)
}

// FromさんがToさんをEventする
// ToさんのconnectsとFromさんもconnectsにお互いを追加する
func (u *EventUseCase) SetFocus(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("EventUseCase.SetFocus Error : from and to is required")
	}
	return u.repo.SetFocus(roomId, from, to)
}

func (u *EventUseCase) DelFocus(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("EventUseCase.DelFocus Error : from and to is required")
	}
	return u.repo.DelFocus(roomId, from, to)
}

func (u *EventUseCase) DelAllFocus(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	if from == "" {
		return fmt.Errorf("EventUseCase.DelAllFocus Error : from is required")
	}
	return u.repo.DelAllFocus(roomId, from)
}

func (u *EventUseCase) GetMemberOfRoomId(roomId entity.RoomId) entity.Members {
	return u.repo.GetMemberOfRoomId(roomId)
}

// CheckExistsEventAndInitはroomIdのeventが登録されているか確認し、登録されてなかった場合eventを初期化して用意します
func (u *EventUseCase) CheckExistsEventAndInit(roomId entity.RoomId) {
	u.repo.CheckExistsEventAndInit(roomId)
}

func (u *EventUseCase) DeleteEventOfRoomId(roomId entity.RoomId) {
	u.repo.DeleteEventOfRoomId(roomId)
}
