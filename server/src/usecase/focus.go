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

func (u *FocusUseCase) SetFocus(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("FocusUseCase.SetFocus Error : from and to is required")
	}
	return u.focusRepo.SetFocus(roomId, from, to)
}
