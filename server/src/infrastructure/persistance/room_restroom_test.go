package persistance

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestRoomRepository_SetMemberRestRoomStateOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		roomId     entity.RoomId
		peerId     entity.PeerId
		isRestRoom bool
		rooms      entity.Rooms
		wantRooms  entity.Rooms
		wantErr    error
	}{
		{
			name:       "正常に動いている場合、渡されたpeerIdのisRestRoomステートを変更する(false->true)",
			roomId:     entity.RoomId("1234"),
			peerId:     entity.PeerId("p1"),
			isRestRoom: true,
			rooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: false,
							},
						},
					},
				},
			},
			wantRooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: true,
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:       "正常に動いている場合、渡されたpeerIdのisRestRoomステートを変更する(true->false)",
			roomId:     entity.RoomId("1234"),
			peerId:     entity.PeerId("p1"),
			isRestRoom: false,
			rooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: true,
							},
						},
					},
				},
			},
			wantRooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: false,
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:       "正常に動いている場合、渡されたpeerIdのisRestRoomステートを変更する(true->true)",
			roomId:     entity.RoomId("1234"),
			peerId:     entity.PeerId("p1"),
			isRestRoom: true,
			rooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: true,
							},
						},
					},
				},
			},
			wantRooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: true,
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:       "正常に動いている場合、渡されたpeerIdのisRestRoomステートを変更する(false->false)",
			roomId:     entity.RoomId("1234"),
			peerId:     entity.PeerId("p1"),
			isRestRoom: false,
			rooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: false,
							},
						},
					},
				},
			},
			wantRooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: false,
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name:       "渡されたpeerIdが指定されたRoomに存在しなかった場合、エラーを返す",
			roomId:     entity.RoomId("1234"),
			peerId:     entity.PeerId("p2"),
			isRestRoom: false,
			rooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: false,
							},
						},
					},
				},
			},
			wantRooms: entity.Rooms{
				entity.RoomId("1234"): &entity.Room{
					MembersStore: &entity.MembersStore{
						Members: entity.Members{
							entity.PeerId("p1"): &entity.Member{
								Name:       entity.Name("hoge"),
								IsRestRoom: false,
							},
						},
					},
				},
			},
			wantErr: fmt.Errorf("RoomRepository.SetMemberRestRoomStateOfRoomId Error : peer_id member not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoom := NewRoomRepository(&tt.rooms)
			err := repoRoom.SetMemberRestRoomStateOfRoomId(tt.roomId, tt.peerId, tt.isRestRoom)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestRoomRepository_SetMemberRestRoomStateOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.wantRooms[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId]) {
				t.Errorf("TestRoomRepository_SetMemberRestRoomStateOfRoomId Error : want %v, but got %v", tt.wantRooms[tt.roomId], (*repoRoom.RoomsStore.Rooms)[tt.roomId])
			}
		})
	}
}
