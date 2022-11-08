package usecase

import (
	"reflect"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestEventUsecase_ExecSoundEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		info      entity.SoundInfo
		wantSound *entity.Sound
		wantErr   error
	}{
		{
			name:   "正常に動いた場合、Soundオブジェクトを返す",
			roomId: entity.RoomId("1234"),
			info: entity.SoundInfo{
				SoundType: entity.SoundType("hoge"),
			},
			wantSound: &entity.Sound{
				SoundType: entity.SoundType("hoge"),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			got, err := ucEvent.ExecSoundEvent(tt.roomId, tt.info)
			if err != tt.wantErr {
				t.Errorf("TestEventUsecase_ExecSoundEvent Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(got, tt.wantSound) {
				t.Errorf("TestEventUsecase_ExecSoundEvent Error : want %v, but got %v", tt.wantSound, got)
			}
		})
	}
}

func TestEventUsecase_isSoundEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		eType entity.EventType
		want  bool
	}{
		{
			name:  "サウンドイベントだった場合、trueが返る",
			eType: entity.SetSound,
			want:  true,
		},
		{
			name:  "サウンドイベント以外の場合、falseが返る",
			eType: entity.SetFocus,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSoundEvent(tt.eType)
			if got != tt.want {
				t.Errorf("TestEventUsecase_isSoundEvent Error : want %v, but got %v", tt.want, got)
			}
		})
	}
}
