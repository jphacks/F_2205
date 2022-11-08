package usecase

import (
	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// ExecEffectEventはEffectMemberにinfoから受け取ったエフェクト情報をいれて返します
func (uc *EventUsecase) ExecEffectEvent(roomId entity.RoomId, info entity.EffectInfo) (*entity.EffectMember, error) {
	m := &entity.EffectMember{
		PeerId: info.PeerId,
		Type:   info.Type,
	}
	return m, nil
}

// isEffectEventは受け取ったイベントがエフェクトイベントか判定します
func isEffectEvent(eType entity.EventType) bool {
	return eType == entity.SetEffect
}
