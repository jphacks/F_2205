package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

var _ IFocusUseCase = &FocusUseCase{}

type FocusUseCase struct {
	focusRepo repository.IFocusRepository
}

type IFocusUseCase interface {
	NewMember(roomId entity.RoomId, info entity.Info) error
	SetFocus(roomId entity.RoomId, info entity.Info) error
	DelFocus(roomId entity.RoomId, info entity.Info) error
	DelAllFocus(roomId entity.RoomId, info entity.Info) error
	GetOrRegisterHub(roomId entity.RoomId) *entity.Hub
	CheckHubExists(roomId entity.RoomId) error
}

func NewFocusUseCase(r repository.IFocusRepository) *FocusUseCase {
	return &FocusUseCase{
		focusRepo: r,
	}
}

func (u *FocusUseCase) NewMember(roomId entity.RoomId, info entity.Info) error {
	newMemberName := info.From
	if newMemberName == "" {
		return fmt.Errorf("FocusUseCase.NewMember Error : from is required")
	}
	return u.focusRepo.NewMember(roomId, newMemberName)
}

// FromさんがToさんをFocusする
// ToさんのconnectsとFromさんもconnectsにお互いを追加する
func (u *FocusUseCase) SetFocus(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("FocusUseCase.SetFocus Error : from and to is required")
	}
	return u.focusRepo.SetFocus(roomId, from, to)
}

func (u *FocusUseCase) DelFocus(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("FocusUseCase.DelFocus Error : from and to is required")
	}
	return u.focusRepo.DelFocus(roomId, from, to)
}

func (u *FocusUseCase) DelAllFocus(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	if from == "" {
		return fmt.Errorf("FocusUseCase.DelAllFocus Error : from is required")
	}
	return u.focusRepo.DelAllFocus(roomId, from)
}

func (u *FocusUseCase) GetOrRegisterHub(roomId entity.RoomId) *entity.Hub {
	return u.focusRepo.GetOrRegisterHub(roomId)
}

func (u *FocusUseCase) CheckHubExists(roomId entity.RoomId) error {
	return u.focusRepo.CheckHubExists(roomId)
}
