package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// SwitchExecFocusEventByEventTypeはフォーカスイベントを実行する関数です
// フォーカスイベントの中の、受け取った特定の処理を実行します
func (uc *EventUsecase) SwitchExecFocusEventByEventType(eType entity.EventType, roomId entity.RoomId, info entity.FocusInfo) error {
	switch eType {
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

// isFocusEventは受け取ったイベントがフォーカスイベントか判定します
func isFocusEvent(eType entity.EventType) bool {
	if eType == entity.SetFocus || eType == entity.DelFocus || eType == entity.DelAllFocus {
		return true
	}
	return false
}
