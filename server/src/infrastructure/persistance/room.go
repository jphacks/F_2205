package persistance

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

var _ repository.IRoomRepository = &RoomRepository{}

type RoomRepository struct {
	Rooms *entity.Rooms
}

func NewRoomRepository(rooms *entity.Rooms) *RoomRepository {
	return &RoomRepository{
		Rooms: rooms,
	}
}

func (r *RoomRepository) AddNewMemberOfRoomId(roomId entity.RoomId, member *entity.Member,peerId entity.PeerId) error {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		return fmt.Errorf("RoomRepository.AddNewMemberOfRoomId Error : room not found")
	}
	_,found = room.Members[peerId]
	if found{
		return fmt.Errorf("RoomRepository.AddNewMemberOfRoomId Error : peer_id already used")
	}
	room.Members[peerId] = member
	return nil
}

// RoomIdが存在するか確認し、存在しなかった場合はroomIdをkeyにRoomを作成する
func (r *RoomRepository) CheckExistsRoomAndInit(roomId entity.RoomId) {
	_, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		r.InitRoomOfRoomId(roomId)
	}
}

// InitRoomOfRoomIdはroomIdをkeyにRoomsに新しく空のRoomオブジェクトを登録します
func (r *RoomRepository) InitRoomOfRoomId(roomId entity.RoomId) {
	(*r.Rooms)[roomId] = &entity.Room{
		Members: entity.Members{},
		FocusMembers: entity.FocusMembers{},
	}
}

// SetNewRoomOfRoomIdはroomIdをkeyにRoomsに新しくRoomを登録します
func (r *RoomRepository) SetNewRoomOfRoomId(room *entity.Room, roomId entity.RoomId) {
	(*r.Rooms)[roomId] = room
}

// DeleteRoomOfRoomIdは指定したRoomIdをKeyとしているRoomをRoomsから削除します
func (r *RoomRepository) DeleteRoomOfRoomId(roomId entity.RoomId) {
	delete(*r.Rooms, roomId)
}

// GetExistsRoomOfRoomIdはroomIdのRoomが存在するか確認し、存在した場合はRoomを返します
func (r *RoomRepository) GetExistsRoomOfRoomId(roomId entity.RoomId) (*entity.Room, bool) {
	room, ok := (*r.Rooms)[roomId]
	if !ok {
		return nil, false
	}
	return room, true
}

func (r *RoomRepository) GetSumOfRoom() int {
	return len(*r.Rooms)
}

// GetMembersOfRoomIdは指定したroomIdのRoomのMemberを返します
// もし見つからなかった場合は空のMembersオブジェクトを返します
func (r *RoomRepository) GetMembersOfRoomId(roomId entity.RoomId) entity.Members {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		// TODO エラーハンドリング
		return entity.Members{}
	}
	return room.Members
}
