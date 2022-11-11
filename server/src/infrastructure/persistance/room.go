package persistance

import (
	"fmt"
	"sync"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
	"github.com/jphacks/F_2205/server/src/domain/service"
)

var _ repository.IRoomRepository = &RoomRepository{}

type RoomRepository struct {
	RoomsStore *entity.RoomsStore
}

// NewRepositoryはrepository.IRoomRepositoryを満たしたRoomRepository構造体を返します
func NewRoomRepository(rooms *entity.Rooms) *RoomRepository {
	return &RoomRepository{
		RoomsStore: &entity.RoomsStore{
			Rooms: rooms,
			Mu:    sync.RWMutex{},
		},
	}
}

// AddNewMemberOfRoomIdは受け取ったRoomIdのRoomに新しくMemberを追加します
func (r *RoomRepository) AddNewMemberOfRoomId(roomId entity.RoomId, member *entity.Member, peerId entity.PeerId) error {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		return fmt.Errorf("RoomRepository.AddNewMemberOfRoomId Error : room not found")
	}
	room.MembersStore.Mu.Lock()
	defer room.MembersStore.Mu.Unlock()

	_, found = room.MembersStore.Members[peerId]
	if found {
		return fmt.Errorf("RoomRepository.AddNewMemberOfRoomId Error : peer_id already used")
	}
	room.MembersStore.Members[peerId] = member
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
	r.RoomsStore.Mu.Lock()
	defer r.RoomsStore.Mu.Unlock()
	(*r.RoomsStore.Rooms)[roomId] = service.NewRoom()
}

// SetNewRoomOfRoomIdはroomIdをkeyにRoomsに新しくRoomを登録します
func (r *RoomRepository) SetNewRoomOfRoomId(room *entity.Room, roomId entity.RoomId) {
	r.RoomsStore.Mu.Lock()
	defer r.RoomsStore.Mu.Unlock()
	(*r.RoomsStore.Rooms)[roomId] = room
}

// DeleteRoomOfRoomIdは指定したRoomIdをKeyとしているRoomをRoomsから削除します
func (r *RoomRepository) DeleteRoomOfRoomId(roomId entity.RoomId) {
	r.RoomsStore.Mu.Lock()
	defer r.RoomsStore.Mu.Unlock()
	delete(*r.RoomsStore.Rooms, roomId)
}

// GetSumOfRoomはサーバーの管理するRoomの数を取得します
func (r *RoomRepository) GetSumOfRoom() int {
	r.RoomsStore.Mu.RLock()
	defer r.RoomsStore.Mu.RUnlock()
	return len(*r.RoomsStore.Rooms)
}

// GetMembersOfRoomIdは指定したroomIdのRoomのMemberを返します
// もし見つからなかった場合は空のMembersオブジェクトを返します
func (r *RoomRepository) GetMembersOfRoomId(roomId entity.RoomId) entity.Members {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		// TODO エラーハンドリング
		return entity.Members{}
	}
	return room.MembersStore.Members
}

// GetExistsRoomOfRoomIdはroomIdのRoomが存在するか確認し、存在した場合はRoomを返します
func (r *RoomRepository) GetExistsRoomOfRoomId(roomId entity.RoomId) (*entity.Room, bool) {
	r.RoomsStore.Mu.RLock()
	defer r.RoomsStore.Mu.RUnlock()
	room, ok := (*r.RoomsStore.Rooms)[roomId]
	if !ok {
		return nil, false
	}
	return room, true
}
