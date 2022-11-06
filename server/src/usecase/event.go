package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

// EventUsecaseはEventのユースケースの構造体です
type EventUsecase struct {
	repoRoom repository.IRoomRepository
}

// IEventUsecaseはEventのユースケースをまとめたインターフェースです
type IEventUsecase interface {
	SwitchExecEventByEventType(e entity.Event, roomId entity.RoomId) (*entity.Room, error)
	SwitchExecFocusEventByEventType(eType entity.EventType, roomId entity.RoomId, info entity.FocusInfo) error
	AddNewMemberOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error
	SetMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error
	DelMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error
	DelAllMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error
	ExecEffectEvent(roomId entity.RoomId, info entity.EffectInfo) (*entity.EffectMember, error)
}

// NewEventUsecaseはIEventUsecaseを満たしたEventUsecase構造体を返します
func NewEventUsecase(repoRoom repository.IRoomRepository) IEventUsecase {
	return &EventUsecase{
		repoRoom: repoRoom,
	}
}

// SwitchExecEventByEventTypeはイベントのタイプによって処理を切り替え、それぞれのイベント実行関数を呼び出します
func (uc *EventUsecase) SwitchExecEventByEventType(e entity.Event, roomId entity.RoomId) (*entity.Room, error) {
	r := &entity.Room{}
	r.EventType = e.Type

	// ScreenShotEventの場合
	if isScreenShotEvent(e.Type) {
		return r, nil
	}

	// FocusEventの場合
	if isFocusEvent(e.Type) {
		err := uc.SwitchExecFocusEventByEventType(e.Type, roomId, e.Info.Focus)
		if err != nil {
			return r, err
		}
		return r, nil
	}

	// EfectEventの場合
	if isEffectEvent(e.Type) {
		m, _ := uc.ExecEffectEvent(roomId, e.Info.Effect)
		r.EffectMember = *m
		return r, nil
	}
	return r, nil
}

// SwitchExecFocusEventByEventTypeはフォーカスイベントを実行する関数です
// フォーカスイベントの中の、受け取った特定の処理を実行します
func (uc *EventUsecase) SwitchExecFocusEventByEventType(eType entity.EventType, roomId entity.RoomId, info entity.FocusInfo) error {
	switch eType {
	case entity.NewMember:
		if err := uc.AddNewMemberOfRoomId(roomId, info); err != nil {
			return err
		}
	case entity.SetFocus:
		if err := uc.SetMemberFocusOfRoomId(roomId, info); err != nil {
			return err
		}

	case entity.DelFocus:
		if err := uc.DelMemberFocusOfRoomId(roomId, info); err != nil {
			return err
		}
	case entity.DelAllFocus:
		if err := uc.DelAllMemberFocusOfRoomId(roomId, info); err != nil {
			return err
		}
	default:
		return fmt.Errorf("EventUsecase.SwitchExecFocusEventByEventType : not matched type")
	}
	return nil
}

// AddNewMemberOfRoomIdは指定されたRoomのFocusMemberに新規メンバーを追加します
func (uc *EventUsecase) AddNewMemberOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error {
	newMemberName := info.From
	if newMemberName == "" {
		return fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : from is required")
	}
	return uc.repoRoom.AddNewMemberOfRoomId(roomId, newMemberName)
}

// SetMemberFocusOfRoomIdは指定されたRoomに新しくフォーカス状態のユーザーを追加します
func (uc *EventUsecase) SetMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("EventUsecase.SetMemberFocusOfRoomId Error : from and to is required")
	}
	return uc.repoRoom.SetMemberFocusOfRoomId(roomId, from, to)
}

// DelMemberFocusOfRoomIdは指定されたRoomの特定のユーザー同士のフォーカスを解除します
func (uc *EventUsecase) DelMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error {
	from := info.From
	to := info.To
	if from == "" || to == "" {
		return fmt.Errorf("EventUsecase.DelMemberFocusOfRoomId Error : from and to is required")
	}
	return uc.repoRoom.DelMemberFocusOfRoomId(roomId, from, to)
}

// DelAllMemberFocusOfRoomIdは指定されたRoomの特定のユーザーのフォーカスをすべて解除します
func (uc *EventUsecase) DelAllMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error {
	from := info.From
	if from == "" {
		return fmt.Errorf("EventUsecase.DelAllMemberFocusOfRoomId Error : from is required")
	}
	return uc.repoRoom.DelAllMemberFocusOfRoomId(roomId, from)
}

// ExecEffectEventはEffectMemberにinfoから受け取ったエフェクト情報をいれて返します
func (uc *EventUsecase) ExecEffectEvent(roomId entity.RoomId, info entity.EffectInfo) (*entity.EffectMember, error) {
	m := &entity.EffectMember{
		Name: info.Name,
		Type: info.Type,
	}
	return m, nil
}

// isScreenShotEventは受け取ったイベントがスクリーンショットイベントか判定します
func isScreenShotEvent(eType entity.EventType) bool {
	return eType == entity.SetScreenShot
}

// isFocusEventは受け取ったイベントがフォーカスイベントか判定します
func isFocusEvent(eType entity.EventType) bool {
	if eType == entity.NewMember || eType == entity.SetFocus ||
		eType == entity.DelFocus || eType == entity.DelAllFocus {
		return true
	}
	return false
}

// isEffectEventは受け取ったイベントがエフェクトイベントか判定します
func isEffectEvent(eType entity.EventType) bool {
	return eType == entity.SetEffect
}
