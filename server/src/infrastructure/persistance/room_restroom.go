package persistance

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)


func (r *RoomRepository) SetMemberRestRoomStateOfRoomId(roomId entity.RoomId,peerId entity.PeerId,isRestRoom bool) error {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		return fmt.Errorf("RoomRepository.SetMemberRestRoomStateOfRoomId Error : room not found")
	}
	member,found := room.Members[peerId]
	if !found{
		return fmt.Errorf("RoomRepository.SetMemberRestRoomStateOfRoomId Error : peer_id member not found")
	}
	member.IsRestRoom = isRestRoom
	return nil
}
