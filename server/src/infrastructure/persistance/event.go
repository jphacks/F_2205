package persistance

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

var _ repository.IEventRepository = &EventRepository{}

type EventRepository struct {
	Events *entity.Events
}

func NewEventRepository(events *entity.Events) *EventRepository {
	return &EventRepository{
		Events: events,
	}
}

func (r *EventRepository) NewMember(roomId entity.RoomId, newMemberName entity.Name) error {
	e, found := r.getExistsEventOfRoomId(roomId)
	if !found {
		return fmt.Errorf("EventRepository.NewMember Error : event not found")
	}

	for _, member := range e.Focus.Members {
		if member.Name == newMemberName {
			return fmt.Errorf("EventRepository.NewMember Error : userName already exist")
		}
	}

	e.Focus.Members = append(
		e.Focus.Members,
		&entity.Member{
			Name:     newMemberName,
			Connects: []*entity.Connect{},
		},
	)
	return nil
}

func (r *EventRepository) SetFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	e, found := r.getExistsEventOfRoomId(roomId)
	if !found {
		return fmt.Errorf("EventRepository.SetFocus Error : event not found")
	}

	// TODO 複数回SetFocusをするとuserが重複してしまう
	for _, member := range e.Focus.Members {
		if member.Name == from {
			for _, connect := range member.Connects {
				if connect != nil && connect.Name == to {
					return fmt.Errorf("EventRepository.SetFocus Error : already connected")
				}
			}

			// FromさんのConnectにToさんを追加
			member.Connects = append(
				member.Connects,
				&entity.Connect{Name: to},
			)
		} else if member.Name == to {
			for _, connect := range member.Connects {
				if connect != nil && connect.Name == to {
					return fmt.Errorf("EventRepository.SetFocus Error : already connected")
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

func (r *EventRepository) DelFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	e, found := r.getExistsEventOfRoomId(roomId)
	if !found {
		return fmt.Errorf("EventRepository.DelFocus Error : event not found")
	}
	for _, member := range e.Focus.Members {
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

func (r *EventRepository) DelAllFocus(roomId entity.RoomId, from entity.Name) error {
	e, found := r.getExistsEventOfRoomId(roomId)
	if !found {
		return fmt.Errorf("EventRepository.DelAllFocus Error : event not found")
	}

	for _, member := range e.Focus.Members {
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

func (r *EventRepository) GetMemberOfRoomId(roomId entity.RoomId) entity.Members {
	e, found := r.getExistsEventOfRoomId(roomId)
	if !found {
		// TODO エラーハンドリング
		return entity.Members{}
	}
	return e.Focus.Members
}

func (r *EventRepository) CheckExistsEventAndInit(roomId entity.RoomId) {
	_, found := r.getExistsEventOfRoomId(roomId)
	if !found {
		r.initEventOfRoomId(roomId)
	}
}

// initEventOfRoomIdはroomIdをkeyにEventsに新しく空のEventオブジェクトを登録します
func (r *EventRepository) initEventOfRoomId(roomId entity.RoomId) {
	(*r.Events)[roomId] = &entity.Event{}
}

// setNewEventOfRoomIdはroomIdをkeyにEventsに新しくEventを登録します
func (r *EventRepository) setNewEventOfRoomId(e *entity.Event, roomId entity.RoomId) {
	(*r.Events)[roomId] = e
}

// deleteEventOfRoomIdは指定したRoomIdをKeyとしているEventをEventsから削除します
func (r *EventRepository) DeleteEventOfRoomId(roomId entity.RoomId) {
	delete(*r.Events, roomId)
}

// getExistsEventOfRoomIdはroomIdのEventが存在するか確認し、存在した場合はEventを返します
func (r *EventRepository) getExistsEventOfRoomId(roomId entity.RoomId) (*entity.Event, bool) {
	e, ok := (*r.Events)[roomId]
	if !ok {
		return nil, false
	}
	return e, true
}
