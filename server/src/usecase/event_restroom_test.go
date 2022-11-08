package usecase

import (
	"fmt"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestEventUsecase_ExecRestRoomEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		roomId  entity.RoomId
		info    entity.RestRoomInfo
		wantErr error
	}{
		{
			name:   "正常に動いた場合",
			roomId: entity.RoomId("1234"),
			info: entity.RestRoomInfo{
				PeerId:     entity.PeerId("p1"),
				IsRestRoom: false,
			},
			wantErr: nil,
		},
		{
			name:   "isRestRoomがなくても、正常に動く",
			roomId: entity.RoomId("1234"),
			info: entity.RestRoomInfo{
				PeerId: entity.PeerId("p1"),
			},
			wantErr: nil,
		},
		{
			name:    "peerIdがない場合、エラーを返す",
			roomId:  entity.RoomId("1234"),
			info:    entity.RestRoomInfo{},
			wantErr: fmt.Errorf("EventUsecase.ExecRestRoomEvent Error : peer_id is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			err := ucEvent.ExecRestRoomEvent(tt.roomId, tt.info)
			if  err!=nil && err.Error()!=tt.wantErr.Error(){
				t.Errorf("TestEventUsecase_ExecRestRoomEvent Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

func TestEventUsecase_isRestRoomEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		eType entity.EventType
		want  bool
	}{
		{
			name:  "お手洗いイベントだった場合、trueが返る",
			eType: entity.SetRestRoom,
			want:  true,
		},
		{
			name:  "お手洗いイベント以外の場合、falseが返る",
			eType: entity.SetFocus,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRestRoomEvent(tt.eType)
			if got != tt.want {
				t.Errorf("TestEventUsecase_isRestRoomEvent Error : want %v, but got %v", tt.want, got)
			}
		})
	}
}
