package usecase

import (
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestEventUsecase_isScreenShotEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		eType entity.EventType
		want  bool
	}{
		{
			name:  "スクリーンショットイベントだった場合、trueが返る",
			eType: entity.SetScreenShot,
			want:  true,
		},
		{
			name:  "スクリーンショットイベント以外の場合、falseが返る",
			eType: entity.SetEffect,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isScreenShotEvent(tt.eType)
			if got != tt.want {
				t.Errorf("TestEventUsecase_isScreenShotEvent Error : want %v, but got %v", tt.want, got)
			}
		})
	}
}
