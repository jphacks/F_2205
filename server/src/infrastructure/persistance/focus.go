package persistance

import (
	"fmt"
	"log"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
)

var _ repository.IFocusRepository = &FocusRepository{}

type FocusRepository struct {
	Hubs *entity.Hubs
}

func NewFocusRepository(hubs *entity.Hubs) *FocusRepository {
	return &FocusRepository{
		Hubs: hubs,
	}
}

func (r *FocusRepository) NewMember(roomId entity.RoomId, newMemberName entity.Name) error {
	h, ok := (*r.Hubs)[roomId]
	if !ok {
		return fmt.Errorf("FocusRepository.NewMember Error : hub not found")
	}

	for _,member := range h.Focus.Members {
		if member.Name == newMemberName {
			return fmt.Errorf("FocusRepository.NewMember Error : userName already exist")
		}
	}

	h.Focus.Members = append(
		h.Focus.Members, 
		&entity.Member{
			Name: newMemberName,
			Connects: []*entity.Connect{},
		},
	)
	log.Println("hub menbers",h.Focus.Members)
	return nil
}

func (r *FocusRepository) SetFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	h, ok := (*r.Hubs)[roomId]
	if !ok {
		return fmt.Errorf("FocusRepository.NewMember Error : hub not found")
	}

	members := h.Focus.Members
	for _,member := range members {
		if member.Name == from{
			// FromさんのConnectにToさんを追加
			member.Connects = append(
				member.Connects, 
				&entity.Connect{Name: to},
			)
		}else if member.Name == to{
			// ToさんのConnectにFromさんを追加
			member.Connects = append(
				member.Connects, 
				&entity.Connect{Name: from},
			)
		}
	}
	return nil
}
