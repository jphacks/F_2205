package persistance

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestRoomRepository_AddNewFocusMemberOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		roomId        entity.RoomId
		newMemberName entity.Name
		room          *entity.Room // 関数実行前のroomの状態
		wantRoom      *entity.Room // 関数実行後の期待するroomの状態
		wantErr       error
	}{
		{
			name:          "正常に動いている場合、新しいメンバーを部屋に追加する",
			roomId:        entity.RoomId("1234"),
			newMemberName: entity.Name("piyo"),
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
					&entity.FocusMember{
						Name:     entity.Name("piyo"),
						Connects: entity.Connects{},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:          "すでに同じ名前の人がいた場合、エラーを返す",
			roomId:        entity.RoomId("1234"),
			newMemberName: entity.Name("hoge"),
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
			wantErr: fmt.Errorf("RoomRepository.AddNewFocusMemberOfRoomId Error : new member name already exist"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// roomMockはroomのmockです
			roomMock := &entity.Rooms{entity.RoomId("1234"): tt.room}
			repoRoom := NewRoomRepository(roomMock)
			err := repoRoom.AddNewFocusMemberOfRoomId(tt.roomId, tt.newMemberName)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomUsecase_AddNewFocusMemberOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(tt.wantRoom, (*repoRoom.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomUsecase_AddNewFocusMemberOfRoomId Error : want %v, but got %v", tt.wantRoom, (*repoRoom.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_SetMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		roomId   entity.RoomId
		from     entity.Name
		to       entity.Name
		room     *entity.Room // 関数実行前のroomの状態
		wantRoom *entity.Room // 関数実行後の期待するroomの状態
		wantErr  error
	}{
		{
			name:   "正常に動いている場合、新しいメンバーを部屋に追加する",
			roomId: entity.RoomId("1234"),
			from:   entity.Name("hoge"),
			to:     entity.Name("piyo"),
			room: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name:     entity.Name("hoge"),
						Connects: entity.Connects{},
					},
					&entity.FocusMember{
						Name:     entity.Name("piyo"),
						Connects: entity.Connects{},
					},
				},
			},
			wantRoom: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name: entity.Name("hoge"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
						},
					},
					&entity.FocusMember{
						Name: entity.Name("piyo"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("hoge"),
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:   "すでにfocus状態だった場合(fromさんのConnectsにtoさんがいた場合)、エラーを返す",
			roomId: entity.RoomId("1234"),
			from:   entity.Name("hoge"),
			to:     entity.Name("piyo"),
			room: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name: entity.Name("hoge"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
						},
					},
					&entity.FocusMember{
						Name:     entity.Name("piyo"),
						Connects: entity.Connects{},
					},
				},
			},
			wantRoom: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name: entity.Name("hoge"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
						},
					},
					&entity.FocusMember{
						Name:     entity.Name("piyo"),
						Connects: entity.Connects{},
					},
				},
			},
			wantErr: fmt.Errorf("RoomRepository.SetFocus Error : already connected"),
		},
		{
			name:   "すでにfocus状態だった場合(toさんのConnectsにfromさんがいた場合)、エラーを返す",
			roomId: entity.RoomId("1234"),
			from:   entity.Name("hoge"),
			to:     entity.Name("piyo"),
			room: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name:     entity.Name("hoge"),
						Connects: entity.Connects{},
					},
					&entity.FocusMember{
						Name: entity.Name("piyo"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("hoge"),
							},
						},
					},
				},
			},
			wantRoom: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name: entity.Name("hoge"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
						},
					},
					&entity.FocusMember{
						Name: entity.Name("piyo"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("hoge"),
							},
						},
					},
				},
			},
			wantErr: fmt.Errorf("RoomRepository.SetFocus Error : already connected"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// roomMockはroomのmockです
			roomMock := &entity.Rooms{
				entity.RoomId("1234"): tt.room,
			}
			repoRoom := NewRoomRepository(roomMock)
			err := repoRoom.SetMemberFocusOfRoomId(tt.roomId, tt.from, tt.to)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomRepository_SetMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(tt.wantRoom, (*repoRoom.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_SetMemberFocusOfRoomId Error : want %v, but got %v", tt.wantRoom, (*repoRoom.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_DelMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		roomId   entity.RoomId
		from     entity.Name
		to       entity.Name
		room     *entity.Room // 関数実行前のroomの状態
		wantRoom *entity.Room // 関数実行後の期待するroomの状態
		wantErr  error
	}{
		{
			name:   "正常に動いている場合、指定されたメンバーを削除する",
			roomId: entity.RoomId("1234"),
			from:   entity.Name("hoge"),
			to:     entity.Name("piyo"),
			room: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name: entity.Name("hoge"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
						},
					},
					&entity.FocusMember{
						Name: entity.Name("piyo"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("hoge"),
							},
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
					&entity.FocusMember{
						Name:     entity.Name("piyo"),
						Connects: entity.Connects{},
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// roomMockはroomのmockです
			roomMock := &entity.Rooms{
				entity.RoomId("1234"): tt.room,
			}
			repoRoom := NewRoomRepository(roomMock)
			err := repoRoom.DelMemberFocusOfRoomId(tt.roomId, tt.from, tt.to)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomRepository_DelMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(tt.wantRoom, (*repoRoom.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_DelMemberFocusOfRoomId Error : want %v, but got %v", tt.wantRoom, (*repoRoom.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_DelAllMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		roomId   entity.RoomId
		from     entity.Name
		room     *entity.Room // 関数実行前のroomの状態
		wantRoom *entity.Room // 関数実行後の期待するroomの状態
		wantErr  error
	}{
		{
			name:   "正常に動いている場合指定されたメンバーを削除する、さらにほかのユーザーとfocus状態だった場合そのfocusも削除する",
			roomId: entity.RoomId("1234"),
			from:   entity.Name("hoge"),
			room: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name: entity.Name("hoge"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
						},
					},
					&entity.FocusMember{
						Name: entity.Name("piyo"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("hoge"),
							},
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
					&entity.FocusMember{
						Name:     entity.Name("piyo"),
						Connects: entity.Connects{},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:   "正常に動いている場合指定されたメンバーを削除する、さらにほかの複数のユーザーとfocus状態だった場合そのfocusも削除する",
			roomId: entity.RoomId("1234"),
			from:   entity.Name("hoge"),
			room: &entity.Room{
				FocusMembers: entity.FocusMembers{
					&entity.FocusMember{
						Name: entity.Name("hoge"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
							&entity.Connect{
								Name: entity.Name("fuga"),
							},
						},
					},
					&entity.FocusMember{
						Name: entity.Name("piyo"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("hoge"),
							},
						},
					},
					&entity.FocusMember{
						Name: entity.Name("fuga"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
							&entity.Connect{
								Name: entity.Name("hoge"),
							},
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
					&entity.FocusMember{
						Name:     entity.Name("piyo"),
						Connects: entity.Connects{},
					},
					&entity.FocusMember{
						Name: entity.Name("fuga"),
						Connects: entity.Connects{
							&entity.Connect{
								Name: entity.Name("piyo"),
							},
						},
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// roomMockはroomのmockです
			// TODO ここが具体的な値である1234に依存しているのはおかしい、なおす
			roomMock := &entity.Rooms{
				entity.RoomId("1234"): tt.room,
			}
			repoRoom := NewRoomRepository(roomMock)
			err := repoRoom.DelAllMemberFocusOfRoomId(tt.roomId, tt.from)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomRepository_DelAllMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(tt.wantRoom, (*repoRoom.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_DelAllMemberFocusOfRoomId Error : want %v, but got %v", tt.wantRoom, (*repoRoom.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_GetFocusMembersOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		roomId              entity.RoomId
		rooms               *entity.Rooms
		wantFocusMembers    entity.FocusMembers
	}{
		{
			name: "正常に動いた場合",
			roomId: entity.RoomId("1234"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"):&entity.Room{
					FocusMembers: entity.FocusMembers{
						&entity.FocusMember{
							Name: entity.Name("hoge"),
							Connects: entity.Connects{},
						},
					},
				},
			},
			wantFocusMembers: entity.FocusMembers{
				&entity.FocusMember{
					Name: entity.Name("hoge"),
					Connects: entity.Connects{},
				},
			},
		},
	}
	for _,tt := range tests {
		t.Run(tt.name,func(t *testing.T) {
			repoRoom := NewRoomRepository(tt.rooms)
			gotFocusMembers := repoRoom.GetFocusMembersOfRoomId(tt.roomId)
			if !reflect.DeepEqual(gotFocusMembers,tt.wantFocusMembers){
				t.Errorf("TestRoomRepository_GetFocusMembersOfRoomId Error : want %v, but got %v", tt.wantFocusMembers, gotFocusMembers)
			}
		})
	}
}
