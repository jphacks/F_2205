package persistance

import (
	"reflect"
	"testing"
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestRoomRepository_AddNewMemberOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		roomId     entity.RoomId
		member     *entity.Member
		peerId     entity.PeerId
		rooms     entity.Rooms
		wantRooms  entity.Rooms
		wantErr    error
	}{
		{
			name:   "正常に動いている場合、Membersに新しいmemberを追加する",
			roomId: entity.RoomId("1234"),
			member: &entity.Member{
				Name: entity.Name("hoge"),
				IsRestRoom: false,
			},
			peerId: entity.PeerId("p1"),
			rooms: entity.Rooms{
				entity.RoomId("1234"):&entity.Room{
					Members:entity.Members{},
				},
			},
			wantRooms: entity.Rooms{
				entity.RoomId("1234"):&entity.Room{
					Members:entity.Members{
						entity.PeerId("p1"):&entity.Member{
							Name: entity.Name("hoge"),
							IsRestRoom: false,
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:   "もしすでに渡されたpeerIdのユーザーが存在していた場合、エラーを返す",
			roomId: entity.RoomId("1234"),
			member: &entity.Member{
				Name: entity.Name("hoge"),
				IsRestRoom: false,
			},
			peerId: entity.PeerId("p1"),
			rooms: entity.Rooms{
				entity.RoomId("1234"):&entity.Room{
					Members:entity.Members{
						entity.PeerId("p1"):&entity.Member{
							Name: entity.Name("hoge"),
							IsRestRoom: false,
						},
					},
				},
			},
			wantRooms: entity.Rooms{
				entity.RoomId("1234"):&entity.Room{
					Members:entity.Members{
						entity.PeerId("p1"):&entity.Member{
							Name: entity.Name("hoge"),
							IsRestRoom: false,
						},
					},
				},
			},
			wantErr: fmt.Errorf("RoomRepository.AddNewMemberOfRoomId Error : peer_id already used"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoom := NewRoomRepository(&tt.rooms)
			err := repoRoom.AddNewMemberOfRoomId(tt.roomId,tt.member,tt.peerId)

			if err !=nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomRepository_AddNewMemberOfRoomId Error : want %v, but got %v", tt.wantErr,err)
			}

			if !reflect.DeepEqual(tt.wantRooms[tt.roomId], (*repoRoom.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_AddNewMemberOfRoomId Error : want %v, but got %v", tt.wantRooms[tt.roomId], (*repoRoom.Rooms)[tt.roomId])
			}
		})
	}
}



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
			name:   "正常に動いている場合、指定されたRoomIdの部屋を生成する",
			roomId: entity.RoomId("1234"),
			wantRoom: &entity.Room{
				Members:      entity.Members{},
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
			room: &entity.Room{},
			wantRoom: &entity.Room{},
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
				entity.RoomId("1234"): &entity.Room{},
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
				entity.RoomId("1234"): &entity.Room{},
			},
			wantRoom: &entity.Room{},
			wantFound: true,
		},
		{
			name:   "指定されたRoomIdの部屋が存在しなければnilが返る",
			roomId: entity.RoomId("4321"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{},
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
				entity.RoomId("1234"): &entity.Room{},
				entity.RoomId("4321"): &entity.Room{},
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

func TestRoomRepository_GetMembersOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		roomId       entity.RoomId
		rooms        *entity.Rooms
		wantMembers  entity.Members
	}{
		{
			name: "正常に動いた場合",
			roomId: entity.RoomId("1234"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"):&entity.Room{
					Members: entity.Members{
						entity.PeerId("p1"):&entity.Member{},
						entity.PeerId("p2"):&entity.Member{},
					},
				},
			},
			wantMembers: entity.Members{
				entity.PeerId("p1"):&entity.Member{},
				entity.PeerId("p2"):&entity.Member{},
			},
		},
		{
			name: "roomIdが存在しなかった場合、空のMembersオブジェクトを返す",
			roomId: entity.RoomId("4321"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"):&entity.Room{
					Members: entity.Members{
						entity.PeerId("p1"):&entity.Member{},
						entity.PeerId("p2"):&entity.Member{},
					},
				},
			},
			wantMembers: entity.Members{},
		},
	}
	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T) {
			roomMock := tt.rooms
			repoRoom := NewRoomRepository(roomMock)
			gotMembers := repoRoom.GetMembersOfRoomId(tt.roomId)
			if !reflect.DeepEqual(gotMembers,tt.wantMembers){
				t.Errorf("TestRoomRepository_GetMembersOfRoomId Error : want %v, but got %v", tt.wantMembers, gotMembers)
			}
		})
	}
}
