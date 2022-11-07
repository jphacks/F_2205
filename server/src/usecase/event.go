package usecase

import (
	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

// EventUsecaseはEventのユースケースの構造体です
type EventUsecase struct {
	repoRoom repository.IRoomRepository
}

// IEventUsecaseはEventのユースケースをまとめたインターフェースです
type IEventUsecase interface {
	// event.go
	SwitchExecEventByEventType(e entity.Event, roomId entity.RoomId) (*entity.Room, error)

	// event_focus.go
	SwitchExecFocusEventByEventType(eType entity.EventType, roomId entity.RoomId, info entity.FocusInfo) error
	AddNewMemberOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error
	SetMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error
	DelMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error
	DelAllMemberFocusOfRoomId(roomId entity.RoomId, info entity.FocusInfo) error

	// event_effect.go
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

	// SoundEventの場合
	if isSoundEvent(e.Type){
		s,_ := uc.ExecSoundEvent(roomId,e.Info.Sound)
		r.Sound = *s
		return r,nil
	}

	return r, nil
}
