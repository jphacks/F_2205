package persistance

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func (r *RoomRepository) AddNewFocusMemberOfRoomId(roomId entity.RoomId, newMemberName entity.Name) error {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		return fmt.Errorf("RoomRepository.AddNewFocusMemberOfRoomId Error : room not found")
	}

	for _, member := range room.FocusMembers {
		if member.Name == newMemberName {
			return fmt.Errorf("RoomRepository.AddNewFocusMemberOfRoomId Error : new member name already exist")
		}
	}

	room.FocusMembers = append(
		room.FocusMembers,
		&entity.FocusMember{
			Name:     newMemberName,
			Connects: entity.Connects{},
		},
	)
	return nil
}

func (r *RoomRepository) SetMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		return fmt.Errorf("RoomRepository.SetMemberFocusOfRoomId Error : room not found")
	}

	for _, member := range room.FocusMembers {
		if member.Name == from {
			for _, connect := range member.Connects {
				if connect != nil && connect.Name == to {
					return fmt.Errorf("RoomRepository.SetFocus Error : already connected")
				}
			}

			// FromさんのConnectにToさんを追加
			member.Connects = append(
				member.Connects,
				&entity.Connect{Name: to},
			)
		} else if member.Name == to {
			for _, connect := range member.Connects {
				if connect != nil && connect.Name == from {
					return fmt.Errorf("RoomRepository.SetFocus Error : already connected")
				}
			}

			// ToさんのConnectにFromさんを追加
			member.Connects = append(
				member.Connects,
				&entity.Connect{Name: from},
			)
		}
	}
	return nil
}

func (r *RoomRepository) DelMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		return fmt.Errorf("RoomRepository.DelMemberFocusOfRoomId Error : room not found")
	}
	for _, member := range room.FocusMembers {
		if member.Name == from {
			// FromさんのConnectからToさんを削除
			for i, connect := range member.Connects {
				if connect != nil && connect.Name == to {
					// 削除
					member.Connects[i] = member.Connects[len(member.Connects)-1]
					member.Connects[len(member.Connects)-1] = nil
					member.Connects = member.Connects[:len(member.Connects)-1]
				}
			}
		} else if member.Name == to {
			// ToさんのConnectからFromさんを削除
			for i, connect := range member.Connects {
				if connect != nil && connect.Name == from {
					// 削除
					member.Connects[i] = member.Connects[len(member.Connects)-1]
					member.Connects[len(member.Connects)-1] = nil
					member.Connects = member.Connects[:len(member.Connects)-1]
				}
			}
		}
	}
	return nil
}

func (r *RoomRepository) DelAllMemberFocusOfRoomId(roomId entity.RoomId, from entity.Name) error {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		return fmt.Errorf("RoomRepository.DelAllMemberFocusOfRoomId Error : room not found")
	}

	for _, member := range room.FocusMembers {
		if member.Name == from {
			// fromさんのConnectをリセット
			member.Connects = []*entity.Connect{}
		} else {
			// fromさん以外の時はfromさんがいないか確認し、あったら削除する
			for i, connect := range member.Connects {
				// TODO こちらのISSUEの対応、なぜnilが入っているのかは調査中
				// https://github.com/jphacks/F_2205/issues/95
				if connect != nil && connect.Name == from {
					// 削除
					member.Connects[i] = member.Connects[len(member.Connects)-1]
					member.Connects[len(member.Connects)-1] = nil
					member.Connects = member.Connects[:len(member.Connects)-1]
				}
			}
		}
	}
	return nil
}

// GetFocusMembersOfRoomIdは指定したroomIdのRoomのFocusMemberを返します
// もし見つからなかった場合は空のFocusMembersオブジェクトを返します
func (r *RoomRepository) GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers {
	room, found := r.GetExistsRoomOfRoomId(roomId)
	if !found {
		// TODO エラーハンドリング
		return entity.FocusMembers{}
	}
	return room.FocusMembers
}
