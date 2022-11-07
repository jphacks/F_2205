package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// ExecRestRoomEventはお手洗いイベントを実行します
func (uc *EventUsecase) ExecRestRoomEvent(roomId entity.RoomId, info entity.RestRoomInfo) error {
	peerId := info.PeerId
	isRestRoom := info.IsRestRoom

	if peerId == "" {
		return fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : peer_id is required")
	}
	if err := uc.repoRoom.SetMemberRestRoomStateOfRoomId(roomId, peerId, isRestRoom); err != nil {
		return fmt.Errorf("EventUsecase.SetMemberRestRoomStateOfRoomId Error : %w", err)
	}
	return nil
}

// isRestRoomEventは受け取ったイベントがお手洗いイベントか判定します
func isRestRoomEvent(eType entity.EventType) bool {
	return eType == entity.SetRestRoom
}
