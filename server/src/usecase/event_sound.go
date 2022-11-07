package usecase

import (
	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// ExecEffectEventはEffectMemberにinfoから受け取ったエフェクト情報をいれて返します
func (uc *EventUsecase) ExecSoundEvent(roomId entity.RoomId, info entity.SoundInfo) (*entity.Sound, error) {
	s := &entity.Sound{
		SoundType: info.SoundType,
	}
	return s, nil
}

// isSoundEventは受け取ったイベントがサウンドイベントか判定します
func isSoundEvent(eType entity.EventType) bool {
	return eType == entity.SetSound
}
