package persistance

import (
	"reflect"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestRoomRepository_CheckExistsRoomAndInit(t *testing.T) {
	t.Skip()
}

func TestRoomRepository_InitRoomOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		roomId   entity.RoomId
		wantRoom *entity.Room
	}{
		{
			name:     "正常に動いている場合、指定されたRoomIdの部屋を生成する",
			roomId:   entity.RoomId("1234"),
			wantRoom: &entity.Room{
				Members: entity.Members{},
				FocusMembers: entity.FocusMembers{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// roomMockはroomのmockです
			roomMock := &entity.Rooms{}
			repoRoom := NewRoomRepository(roomMock)
			repoRoom.InitRoomOfRoomId(tt.roomId)
			if !reflect.DeepEqual(tt.wantRoom, (*repoRoom.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_InitRoomOfRoomId Error : want %v, but got %v", tt.wantRoom, (*repoRoom.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_SetNewRoomOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		roomId   entity.RoomId
		room     *entity.Room
		wantRoom *entity.Room
	}{
		{
			name:   "正常に動いている場合、指定されたRoomIdの部屋を生成する",
			roomId: entity.RoomId("1234"),
			room: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name:     entity.Name("hoge"),
						Connects: entity.Connects{},
					},
				},
			},
			wantRoom: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name:     entity.Name("hoge"),
						Connects: entity.Connects{},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roomMock := &entity.Rooms{}
			repoRoom := NewRoomRepository(roomMock)
			repoRoom.SetNewRoomOfRoomId(tt.room, tt.roomId)
			if !reflect.DeepEqual(tt.wantRoom, (*repoRoom.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_SetNewRoomOfRoomId Error : want %v, but got %v", tt.wantRoom, (*repoRoom.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_DeleteRoomOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		rooms     *entity.Rooms
		wantRooms *entity.Rooms
	}{
		{
			name:   "正常に動いている場合、指定されたRoomIdの部屋を削除する",
			roomId: entity.RoomId("1234"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembers: entity.FocusMembers{
						&entity.FocusMember{
							Name:     entity.Name("hoge"),
							Connects: entity.Connects{},
						},
					},
				},
			},
			wantRooms: &entity.Rooms{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roomMock := tt.rooms
			repoRoom := NewRoomRepository(roomMock)
			repoRoom.DeleteRoomOfRoomId(tt.roomId)
			if !reflect.DeepEqual(tt.wantRooms, repoRoom.Rooms) {
				t.Errorf("TestRoomRepository_DeleteRoomOfRoomId Error : want %v, but got %v", tt.wantRooms, repoRoom.Rooms)
			}
		})
	}
}

func TestRoomRepository_GetExistsRoomOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		rooms     *entity.Rooms
		wantRoom  *entity.Room
		wantFound bool
	}{
		{
			name:   "正常に動いている場合、指定されたRoomIdの部屋を取得する",
			roomId: entity.RoomId("1234"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembers: entity.FocusMembers{
						&entity.FocusMember{
							Name:     entity.Name("hoge"),
							Connects: entity.Connects{},
						},
					},
				},
			},
			wantRoom: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name:     entity.Name("hoge"),
						Connects: entity.Connects{},
					},
				},
			},
			wantFound: true,
		},
		{
			name:   "指定されたRoomIdの部屋が存在しなければnilが返る",
			roomId: entity.RoomId("4321"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembers: entity.FocusMembers{
						&entity.FocusMember{
							Name:     entity.Name("hoge"),
							Connects: entity.Connects{},
						},
					},
				},
			},
			wantRoom:  nil,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roomMock := tt.rooms
			repoRoom := NewRoomRepository(roomMock)
			gotRoom, gotFound := repoRoom.GetExistsRoomOfRoomId(tt.roomId)
			if !reflect.DeepEqual(tt.wantRoom, gotRoom) {
				t.Errorf("TestRoomRepository_GetExistsRoomOfRoomId Error : want %v, but got %v", tt.wantRoom, gotRoom)
			}
			if tt.wantFound != gotFound {
				t.Errorf("TestRoomRepository_GetExistsRoomOfRoomId Error : want %v, but got %v", tt.wantFound, gotFound)
			}
		})
	}
}

func TestRoomRepository_GetSumOfRoom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		rooms       *entity.Rooms
		wantRoomLen int
	}{
		{
			name: "正常に動いている場合、Roomの数を返す",
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembers: entity.FocusMembers{
						&entity.FocusMember{
							Name:     entity.Name("hoge"),
							Connects: entity.Connects{},
						},
					},
				},
				entity.RoomId("4321"): &entity.Room{
					FocusMembers: entity.FocusMembers{
						&entity.FocusMember{
							Name:     entity.Name("hoge"),
							Connects: entity.Connects{},
						},
					},
				},
			},
			wantRoomLen: 2,
		},
		{
			name:        "Roomが0個の時も正常に返す",
			rooms:       &entity.Rooms{},
			wantRoomLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoom := NewRoomRepository(tt.rooms)
			gotRoomLen := repoRoom.GetSumOfRoom()
			if tt.wantRoomLen != gotRoomLen {
				t.Errorf("TestRoomRepository_GetSumOfRoom Error : want %v, but got %v", tt.wantRoomLen, gotRoomLen)
			}
		})
	}
}
