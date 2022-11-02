package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

var _ IRoomUseCase = &RoomUseCase{}

type RoomUseCase struct {
	repo repository.IRoomRepository
}

type IRoomUseCase interface {
	ExecEventOfEventType(eType entity.EventType, roomId entity.RoomId, info entity.Info) error
	GetMemberOfRoomId(roomId entity.RoomId) entity.Members
	CheckExistsRoomAndInit(roomId entity.RoomId)
	DeleteRoomOfRoomId(roomId entity.RoomId)
}

func NewRoomUseCase(repo repository.IRoomRepository) *RoomUseCase {
	return &RoomUseCase{
		repo: repo,
	}
}

func (uc *RoomUseCase) GetMemberOfRoomId(roomId entity.RoomId) entity.Members {
	return uc.repo.GetMemberOfRoomId(roomId)
}

// CheckExistsRoomAndInitはroomIdのRoomが登録されているか確認し、登録されてなかった場合Roomを初期化して用意します
func (uc *RoomUseCase) CheckExistsRoomAndInit(roomId entity.RoomId) {
	uc.repo.CheckExistsRoomAndInit(roomId)
}

func (uc *RoomUseCase) DeleteRoomOfRoomId(roomId entity.RoomId) {
	uc.repo.DeleteRoomOfRoomId(roomId)
}

// ExecEventOfEventTypeは指定されたEventTypeのEventを実行します
func (uc *RoomUseCase) ExecEventOfEventType(eType entity.EventType, roomId entity.RoomId, info entity.Info) error {
	switch eType {
	case entity.NewMember:
		if err := uc.addNewMemberOfRoomId(roomId, info); err != nil {
			return err
		}
	case entity.SetFocus:
		if err := uc.setMemberFocusOfRoomId(roomId, info); err != nil {
			return err
		}
	case entity.DelFocus:
		if err := uc.delMemberFocusOfRoomId(roomId, info); err != nil {
			return err
		}
	case entity.DelAllFocus:
		if err := uc.delAllMemberFocusOfRoomId(roomId, info); err != nil {
			return err
		}
	case entity.SetScreenShot:
		// SetScreenShotの場合、何もせず、eventTypeをBroadcastに設定する
	default:
		return fmt.Errorf("RoomUseCase.ExecEventOfEventType : not matched type")
	}
	return nil
}

func (uc *RoomUseCase) addNewMemberOfRoomId(roomId entity.RoomId, info entity.Info) error {
	newMemberName := info.From
	if newMemberName == "" {
		return fmt.Errorf("RoomUseCase.AddNewMemberOfRoomId Error : from is required")
	}
	return uc.repo.AddNewMemberOfRoomId(roomId, newMemberName)
}

// FromさんがToさんをRoomする
// ToさんのconnectsとFromさんもconnectsにお互いを追加する
func (uc *RoomUseCase) setMemberFocusOfRoomId(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("RoomUseCase.SetMemberFocusOfRoomId Error : from and to is required")
	}
	return uc.repo.SetMemberFocusOfRoomId(roomId, from, to)
}

func (uc *RoomUseCase) delMemberFocusOfRoomId(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("RoomUseCase.DelMemberFocusOfRoomId Error : from and to is required")
	}
	return uc.repo.DelMemberFocusOfRoomId(roomId, from, to)
}

func (uc *RoomUseCase) delAllMemberFocusOfRoomId(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	if from == "" {
		return fmt.Errorf("RoomUseCase.DelAllMemberFocusOfRoomId Error : from is required")
	}
	return uc.repo.DelAllMemberFocusOfRoomId(roomId, from)
}

func (uc *RoomUseCase) setScreenShotOfRoomId(roomId entity.RoomId, info entity.Info) error {
	from := info.From
	if from == "" {
		return fmt.Errorf("RoomUseCase.DelAllMemberFocusOfRoomId Error : from is required")
	}
	return uc.repo.DelAllMemberFocusOfRoomId(roomId, from)
}
