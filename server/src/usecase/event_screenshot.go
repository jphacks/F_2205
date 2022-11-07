package usecase

import (
	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// isScreenShotEventは受け取ったイベントがスクリーンショットイベントか判定します
func isScreenShotEvent(eType entity.EventType) bool {
	return eType == entity.SetScreenShot
}
