package persistance

import (
	"fmt"

	"reflect"
	"sync"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestRoomRepository_AddNewFocusMemberOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		roomId          entity.RoomId
		newMemberPeerId entity.PeerId
		rooms           *entity.Rooms // 関数実行前のroomsの状態
		wantRooms       *entity.Rooms // 関数実行後の期待するroomsの状態
		wantErr         error
	}{
		{
			name:            "正常に動いている場合、新しいメンバーを部屋に追加する",
			roomId:          entity.RoomId("1234"),
			newMemberPeerId: entity.PeerId("p2"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId:   entity.PeerId("p2"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:            "すでに同じ名前の人がいた場合、エラーを返す",
			roomId:          entity.RoomId("1234"),
			newMemberPeerId: entity.PeerId("p1"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantErr: fmt.Errorf("RoomRepository.AddNewFocusMemberOfRoomId Error : new member name already exist"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoom := NewRoomRepository(tt.rooms)
			err := repoRoom.AddNewFocusMemberOfRoomId(tt.roomId, tt.newMemberPeerId)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomUsecase_AddNewFocusMemberOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual((*tt.wantRooms)[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomUsecase_AddNewFocusMemberOfRoomId Error : want %v, but got %v", (*tt.wantRooms)[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_SetMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		from      entity.PeerId
		to        entity.PeerId
		rooms     *entity.Rooms // 関数実行前のroomsの状態
		wantRooms *entity.Rooms // 関数実行後の期待するroomsの状態
		wantErr   error
	}{
		{
			name:   "正常に動いている場合、新しいメンバーを部屋に追加する",
			roomId: entity.RoomId("1234"),
			from:   entity.PeerId("p1"),
			to:     entity.PeerId("p2"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId:   entity.PeerId("p2"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId: entity.PeerId("p1"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
								},
								Mu: sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId: entity.PeerId("p2"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p1"),
									},
								},
								Mu: sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:   "すでにfocus状態だった場合(fromさんのConnectsにtoさんがいた場合)、エラーを返す",
			roomId: entity.RoomId("1234"),
			from:   entity.PeerId("p1"),
			to:     entity.PeerId("p2"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId: entity.PeerId("p1"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
								},
								Mu: sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId:   entity.PeerId("p2"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId: entity.PeerId("p1"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
								},
								Mu: sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId:   entity.PeerId("p2"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantErr: fmt.Errorf("RoomRepository.SetFocus Error : already connected"),
		},
		{
			name:   "すでにfocus状態だった場合(toさんのConnectsにfromさんがいた場合)、エラーを返す",
			roomId: entity.RoomId("1234"),
			from:   entity.PeerId("p1"),
			to:     entity.PeerId("p2"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId: entity.PeerId("p2"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p1"),
									},
								},
								Mu: sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId: entity.PeerId("p1"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
								},
								Mu: sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId: entity.PeerId("p2"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p1"),
									},
								},
								Mu: sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantErr: fmt.Errorf("RoomRepository.SetFocus Error : already connected"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoom := NewRoomRepository(tt.rooms)
			err := repoRoom.SetMemberFocusOfRoomId(tt.roomId, tt.from, tt.to)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomRepository_SetMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual((*tt.wantRooms)[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_SetMemberFocusOfRoomId Error : want %v, but got %v", (*tt.wantRooms)[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_DelMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		from      entity.PeerId
		to        entity.PeerId
		rooms     *entity.Rooms // 関数実行前のroomsの状態
		wantRooms *entity.Rooms // 関数実行後の期待するroomsの状態
		wantErr   error
	}{
		{
			name:   "正常に動いている場合、指定されたメンバーを削除する",
			roomId: entity.RoomId("1234"),
			from:   entity.PeerId("p1"),
			to:     entity.PeerId("p2"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId: entity.PeerId("p1"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
								},
								Mu: sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId: entity.PeerId("p2"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p1"),
									},
								},
								Mu: sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId:   entity.PeerId("p2"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoom := NewRoomRepository(tt.rooms)
			err := repoRoom.DelMemberFocusOfRoomId(tt.roomId, tt.from, tt.to)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomRepository_DelMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual((*tt.wantRooms)[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_DelMemberFocusOfRoomId Error : want %v, but got %v", (*tt.wantRooms)[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_DelAllMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		from      entity.PeerId
		rooms     *entity.Rooms // 関数実行前のroomsの状態
		wantRooms *entity.Rooms // 関数実行後の期待するroomsの状態
		wantErr   error
	}{
		{
			name:   "正常に動いている場合指定されたメンバーを削除する、さらにほかのユーザーとfocus状態だった場合そのfocusも削除する",
			roomId: entity.RoomId("1234"),
			from:   entity.PeerId("p1"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId: entity.PeerId("p1"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
								},
								Mu: sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId: entity.PeerId("p2"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p1"),
									},
								},
								Mu: sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId:   entity.PeerId("p2"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:   "正常に動いている場合指定されたメンバーを削除する、さらにほかの複数のユーザーとfocus状態だった場合そのfocusも削除する",
			roomId: entity.RoomId("1234"),
			from:   entity.PeerId("p1"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId: entity.PeerId("p1"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
									&entity.Connect{
										PeerId: entity.PeerId("p3"),
									},
								},
								Mu: sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId: entity.PeerId("p2"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p1"),
									},
								},
								Mu: sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId: entity.PeerId("p3"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
									&entity.Connect{
										PeerId: entity.PeerId("p1"),
									},
								},
								Mu: sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId:   entity.PeerId("p2"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
							&entity.FocusMember{
								PeerId: entity.PeerId("p3"),
								Connects: entity.Connects{
									&entity.Connect{
										PeerId: entity.PeerId("p2"),
									},
								},
								Mu: sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoom := NewRoomRepository(tt.rooms)
			err := repoRoom.DelAllMemberFocusOfRoomId(tt.roomId, tt.from)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomRepository_DelAllMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual((*tt.wantRooms)[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_DelAllMemberFocusOfRoomId Error : want %v, but got %v", (*tt.wantRooms)[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId])
			}
		})
	}
}

func TestRoomRepository_GetFocusMembersOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		rooms     *entity.Rooms
		wantRooms *entity.Rooms
	}{
		{
			name:   "正常に動いた場合",
			roomId: entity.RoomId("1234"),
			rooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
			wantRooms: &entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					FocusMembersStore: &entity.FocusMembersStore{
						FocusMembers: entity.FocusMembers{
							&entity.FocusMember{
								PeerId:   entity.PeerId("p1"),
								Connects: entity.Connects{},
								Mu:       sync.RWMutex{},
							},
						},
						Mu: sync.RWMutex{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoom := NewRoomRepository(tt.rooms)
			gotFocusMembers := repoRoom.GetFocusMembersOfRoomId(tt.roomId)
			if !reflect.DeepEqual(gotFocusMembers, (*tt.wantRooms)[tt.roomId].FocusMembersStore.FocusMembers) {
				t.Errorf("TestRoomRepository_GetFocusMembersOfRoomId Error : want %v, but got %v", (*tt.wantRooms)[tt.roomId].FocusMembersStore.FocusMembers, gotFocusMembers)
			}
		})
	}
}
