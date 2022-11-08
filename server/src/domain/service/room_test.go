package service

import (
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestRoomService_StringToRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct{
		name         string
		roomIdString string
		wantRoomId   entity.RoomId
	}{
		{
			name: "正常に動いた場合、stringからentity.RoomIdにキャストされる",
			roomIdString: "1111",
			wantRoomId: entity.RoomId("1111"),
		},
	}
	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T) {
			got := StringToRoomId(tt.roomIdString)
			if tt.wantRoomId != got {
				t.Errorf("TestRoomService_StringToRoomId Error : want %v, but got %v", tt.wantRoomId,got)
			}
		})
	}
}
