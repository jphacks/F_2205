package usecase

import (
	"testing"
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

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
				Name: entity.Name("hoge"),
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
			name:   "両方ない場合、エラーを返す",
			roomId: entity.RoomId("1234"),
			memberInfo: entity.MemberInfo{},
			wantErr: fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : name is required"),
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