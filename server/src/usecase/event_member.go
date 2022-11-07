package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// SwitchExecMemberEventByEventTypeは受け取ったイベントのタイプで実行する処理を切り替えます
func (uc *EventUsecase) SwitchExecMemberEventByEventType(eType entity.EventType, roomId entity.RoomId, info entity.MemberInfo) error {
	switch eType {
	case entity.AddNewMember:
		if err := uc.AddNewMemberOfRoomId(roomId, info); err != nil {
			return err
		}
	default:
		return fmt.Errorf("EventUsecase.SwitchExecFocusEventByEventType : not matched type")
	}
	return nil
}

// AddNewMemberOfRoomIdは指定されたRoomのMember,FocusMemberに新規メンバーを追加します
func (uc *EventUsecase) AddNewMemberOfRoomId(roomId entity.RoomId, info entity.MemberInfo) error {
	name := info.Name
	peerId := info.PeerId

	if name == "" {
		return fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : name is required")
	}
	if peerId == "" {
		return fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : peer_id is required")
	}
	member := &entity.Member{
		Name:       name,
		IsRestRoom: false,
	}
	if err := uc.repoRoom.AddNewMemberOfRoomId(roomId, member, peerId); err != nil {
		return fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : %w", err)
	}
	if err := uc.repoRoom.AddNewFocusMemberOfRoomId(roomId, member.Name); err != nil {
		return fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : %w", err)
	}
	return nil
}

// isMemberEventは受け取ったイベントがメンバーイベントか判定します
func isMemberEvent(eType entity.EventType) bool {
	return eType == entity.AddNewMember
}
