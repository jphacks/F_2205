package usecase

import (
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestRoomUsecase_CreateRoomNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		wantInfo *entity.RoomInfo
		wantErr  error
	}{
		{
			name: "正常に動いている場合",
			wantInfo: &entity.RoomInfo{
				Id: "1234",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucRoom := NewRoomUsecase(&roomRepositoryMock{})
			got, err := ucRoom.CreateRoomNumber()
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomUsecase_CreateRoomNumber Error : want %v, but got %v", tt.wantErr, err)
			}
			if len(tt.wantInfo.Id) != len(got.Id) {
				t.Errorf("TestRoomUsecase_CreateRoomNumber Error : want %v, but got %v", len(tt.wantInfo.Id), len(got.Id))
			}
		})
	}
}

// roomRepositoryMockはroomRepositoryのmockの構造体です
type roomRepositoryMock struct{}

func (r *roomRepositoryMock) AddNewMemberOfRoomId(roomId entity.RoomId, newMemberName entity.Name) error {
	return nil
}
func (r *roomRepositoryMock) SetMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	return nil
}
func (r *roomRepositoryMock) DelMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	return nil
}
func (r *roomRepositoryMock) DelAllMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name) error {
	return nil
}
func (r *roomRepositoryMock) GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers {
	return entity.FocusMembers{}
}
func (r *roomRepositoryMock) CheckExistsRoomAndInit(roomId entity.RoomId) {}
func (r *roomRepositoryMock) DeleteRoomOfRoomId(roomId entity.RoomId)     {}
func (r *roomRepositoryMock) GetExistsRoomOfRoomId(roomId entity.RoomId) (*entity.Room, bool) {
	return &entity.Room{}, false
}
func (r *roomRepositoryMock) GetSumOfRoom() int { return 0 }
