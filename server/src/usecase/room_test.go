package usecase

import (
	"fmt"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

// TestRoomUsecase_CreateRoomはusecaseのCreateRoomをテストします
func TestRoomUsecase_CreateRoom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		room    *entity.Room
		wantErr error
	}{
		{
			name: "正常に動いている場合",
			room: &entity.Room{
				Name: "hoge",
			},
			wantErr: nil,
		},
		{
			name:    "必要なフィールド(name)がなかった場合",
			room:    &entity.Room{},
			wantErr: fmt.Errorf("RoomUseCase.CreateRoom Name Error : room name required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roomUC := NewRoomUseCase(&roomRepositoryMock{})
			_, err := roomUC.CreateRoom(tt.room)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomUsecase_GetAllRoom Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

// roomRepositoryMockはroomRepositoryのmockの構造体です
type roomRepositoryMock struct{}

func (r *roomRepositoryMock) CreateRoom(room *entity.Room) (*entity.Room, error) {
	return room, nil
}
