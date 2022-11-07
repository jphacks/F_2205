package usecase

import (
	"reflect"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestEventUsecase_ExecEffectEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		roomId           entity.RoomId
		effectInfo       entity.EffectInfo
		wantEffectMember *entity.EffectMember
		wantErr          error
	}{
		{
			name:   "正常に動いている場合",
			roomId: entity.RoomId("1234"),
			effectInfo: entity.EffectInfo{
				Name: entity.Name("hoge"),
				Type: entity.EffectType("happy"),
			},
			wantEffectMember: &entity.EffectMember{
				Name: entity.Name("hoge"),
				Type: entity.EffectType("happy"),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			got, err := ucEvent.ExecEffectEvent(tt.roomId, tt.effectInfo)
			if !reflect.DeepEqual(tt.wantEffectMember, got) {
				t.Errorf("TestEventUsecase_ExecEffectEvent Error : want %v, but got %v", tt.wantEffectMember, got)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestEventUsecase_ExecEffectEvent Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

func TestEventUsecase_isEffectEvent(t *testing.T){
	t.Parallel()

	tests := []struct{
		Name string
		eType entity.EventType
		want bool
	}{
		{
			Name: "エフェクトイベントだった場合、trueが返る",
			eType: entity.SetEffect,
			want: true,
		},
		{
			Name: "エフェクトイベント以外の場合、falseが返る",
			eType: entity.SetFocus,
			want: false,
		},
	}
	for _,tt := range tests {
		t.Run(tt.Name,func(t *testing.T) {
			got := isEffectEvent(tt.eType)
			if got!=tt.want {
				t.Errorf("TestEventUsecase_isEffectEvent Error : want %v, but got %v", tt.want, got)
			}
		})
	}
}
