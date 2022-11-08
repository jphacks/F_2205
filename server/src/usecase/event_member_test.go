package usecase

import (
	"fmt"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestEventUsecase_SwitchExecMemberEventByEventType(t *testing.T) {
	t.Skip()
}

func TestEventUsecase_AddNewMemberOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		roomId     entity.RoomId
		memberInfo entity.MemberInfo
		wantRoom   *entity.Room
		wantErr    error
	}{
		{
			name:   "正常に動いている場合",
			roomId: entity.RoomId("1234"),
			memberInfo: entity.MemberInfo{
				PeerId: "hoge",
				Name:   entity.Name("hoge"),
			},
			wantErr: nil,
		},
		{
			name:   "peer_idがない場合、エラーを返す",
			roomId: entity.RoomId("1234"),
			memberInfo: entity.MemberInfo{
				Name: entity.Name("hoge"),
			},
			wantErr: fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : peer_id is required"),
		},
		{
			name:   "nameがない場合、エラーを返す",
			roomId: entity.RoomId("1234"),
			memberInfo: entity.MemberInfo{
				PeerId: "hoge",
			},
			wantErr: fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : name is required"),
		},
		{
			name:       "両方ない場合、エラーを返す",
			roomId:     entity.RoomId("1234"),
			memberInfo: entity.MemberInfo{},
			wantErr:    fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : name is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			err := ucEvent.AddNewMemberOfRoomId(tt.roomId, tt.memberInfo)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestEventUsecase_AddNewMemberOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}


func TestEventUsecase_isMemberEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name  string
		eType entity.EventType
		want  bool
	}{
		{
			Name:  "メンバーイベントだった場合、trueが返る",
			eType: entity.AddNewMember,
			want:  true,
		},
		{
			Name:  "メンバーイベント以外の場合、falseが返る",
			eType: entity.SetFocus,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := isMemberEvent(tt.eType)
			if got != tt.want {
				t.Errorf("TestEventUsecase_isMemberEvent Error : want %v, but got %v", tt.want, got)
			}
		})
	}
}
