package usecase

import (
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestEventUsecase_isSoundEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name  string
		eType entity.EventType
		want  bool
	}{
		{
			Name:  "サウンドイベントだった場合、trueが返る",
			eType: entity.SetSound,
			want:  true,
		},
		{
			Name:  "サウンドイベント以外の場合、falseが返る",
			eType: entity.SetFocus,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := isSoundEvent(tt.eType)
			if got != tt.want {
				t.Errorf("TestEventUsecase_isSoundEvent Error : want %v, but got %v", tt.want, got)
			}
		})
	}
}
