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
	ExecEventOfEventType(e entity.Event, roomId entity.RoomId) (*entity.Room, error)
	GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers
	CheckExistsRoomAndInit(roomId entity.RoomId)
	DeleteRoomOfRoomId(roomId entity.RoomId)
}

func NewRoomUseCase(repo repository.IRoomRepository) *RoomUseCase {
	return &RoomUseCase{
		repo: repo,
	}
}

func (uc *RoomUseCase) GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers {
	return uc.repo.GetFocusMembersOfRoomId(roomId)
}

// CheckExistsRoomAndInitはroomIdのRoomが登録されているか確認し、登録されてなかった場合Roomを初期化して用意します
func (uc *RoomUseCase) CheckExistsRoomAndInit(roomId entity.RoomId) {
	uc.repo.CheckExistsRoomAndInit(roomId)
}

func (uc *RoomUseCase) DeleteRoomOfRoomId(roomId entity.RoomId) {
	uc.repo.DeleteRoomOfRoomId(roomId)
}

// ExecEventOfEventTypeはInfoの型をチェックし、それぞれのイベント実行関数を呼び出します
func (uc *RoomUseCase) ExecEventOfEventType(e entity.Event, roomId entity.RoomId) (*entity.Room, error) {
	r := &entity.Room{}
	r.EventType = e.Type

	// ScreenShotEventの場合
	if e.Type == entity.SetScreenShot {
		return r, nil
	}

	// FocusEventの場合
	// TODO もう少しきれいに書きたい
	if e.Type == entity.NewMember || e.Type == entity.SetFocus ||
		e.Type == entity.DelFocus || e.Type == entity.DelAllFocus {
		// FocusEventを実行する
		err := uc.execFocusEventOfEventType(e.Type, roomId, e.Info.Focus)
		if err != nil {
			return r, err
		}
		return r, nil
	}

	// EfectEventの場合
	if e.Type == entity.SetEffect {
		m, _ := uc.execEffectEventOfEventType(roomId, e.Info.Effect)
		r.EffectMember = *m
		return r, nil
	}
	return r, nil
}

// execFocusEventOfEventTypeは指定されたEventTypeのEventを実行します
func (uc *RoomUseCase) execFocusEventOfEventType(
	eType entity.EventType,
	roomId entity.RoomId,
	info entity.FocusInfo) error {
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
	default:
		return fmt.Errorf("RoomUseCase.ExecEventOfEventType : not matched type")
	}
	return nil
}

func (uc *RoomUseCase) execEffectEventOfEventType(roomId entity.RoomId, info entity.EffectInfo) (*entity.EffectMember, error) {
	m := &entity.EffectMember{
		Name: info.Name,
		Type: info.Type,
	}
	return m, nil
}

func (uc *RoomUseCase) addNewMemberOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error {
	newMemberName := info.From
	if newMemberName == "" {
		return fmt.Errorf("RoomUseCase.AddNewMemberOfRoomId Error : from is required")
	}
	return uc.repo.AddNewMemberOfRoomId(roomId, newMemberName)
}

// FromさんがToさんをRoomする
// ToさんのconnectsとFromさんもconnectsにお互いを追加する
func (uc *RoomUseCase) setMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("RoomUseCase.SetMemberFocusOfRoomId Error : from and to is required")
	}
	return uc.repo.SetMemberFocusOfRoomId(roomId, from, to)
}

func (uc *RoomUseCase) delMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("RoomUseCase.DelMemberFocusOfRoomId Error : from and to is required")
	}
	return uc.repo.DelMemberFocusOfRoomId(roomId, from, to)
}

func (uc *RoomUseCase) delAllMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error {
	from := info.From
	if from == "" {
		return fmt.Errorf("RoomUseCase.DelAllMemberFocusOfRoomId Error : from is required")
	}
	return uc.repo.DelAllMemberFocusOfRoomId(roomId, from)
}
