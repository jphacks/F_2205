package persistance

import (
	"fmt"
	"log"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
	"github.com/jphacks/F_2205/server/src/infrastructure/hub"
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
	// TODO serviceとかにしてもいいかも
	h, ok := (*r.Hubs)[roomId]
	if !ok {
		return fmt.Errorf("FocusRepository.NewMember Error : hub not found")
	}

	for _, member := range h.Focus.Members {
		if member.Name == newMemberName {
			return fmt.Errorf("FocusRepository.NewMember Error : userName already exist")
		}
	}

	h.Focus.Members = append(
		h.Focus.Members,
		&entity.Member{
			Name:     newMemberName,
			Connects: []*entity.Connect{},
		},
	)
	log.Println("hub menbers", h.Focus.Members)
	return nil
}

func (r *FocusRepository) SetFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	h, ok := (*r.Hubs)[roomId]
	if !ok {
		return fmt.Errorf("FocusRepository.SetFocus Error : hub not found")
	}

	for _, member := range h.Focus.Members {
		if member.Name == from {
			for _, connect := range member.Connects {
				if connect.Name == to {
					return fmt.Errorf("FocusRepository.SetFocus Error : already connected")
				}
			}

			// FromさんのConnectにToさんを追加
			member.Connects = append(
				member.Connects,
				&entity.Connect{Name: to},
			)
		} else if member.Name == to {
			for _, connect := range member.Connects {
				if connect.Name == to {
					return fmt.Errorf("FocusRepository.SetFocus Error : already connected")
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

func (r *FocusRepository) DelFocus(roomId entity.RoomId, from entity.Name, to entity.Name) error {
	h, ok := (*r.Hubs)[roomId]
	if !ok {
		return fmt.Errorf("FocusRepository.DelFocus Error : hub not found")
	}
	for _, member := range h.Focus.Members {
		if member.Name == from {
			// FromさんのConnectからToさんを削除
			for i, connect := range member.Connects {
				if connect.Name == to {
					// 削除
					member.Connects[i] = member.Connects[len(member.Connects)-1]
					member.Connects[len(member.Connects)-1] = nil
					member.Connects = member.Connects[:len(member.Connects)-1]
				}
			}
		} else if member.Name == to {
			// ToさんのConnectからFromさんを削除
			for i, connect := range member.Connects {
				if connect.Name == from {
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

func (r *FocusRepository) DelAllFocus(roomId entity.RoomId, from entity.Name) error {
	h, ok := (*r.Hubs)[roomId]
	if !ok {
		return fmt.Errorf("FocusRepository.DelAllFocus Error : hub not found")
	}

	for _, member := range h.Focus.Members {
		if member.Name == from {
			// fromさんのConnectをリセット
			member.Connects =  []*entity.Connect{}
		}else{
			// fromさん以外の時はfromさんがいないか確認し、あったら削除する
			for i, connect := range member.Connects {
				if connect.Name == from {
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

func (r *FocusRepository) GetOrRegisterHub(roomId entity.RoomId) *entity.Hub {
	var h *entity.Hub

	if found, ok := (*r.Hubs)[roomId]; ok {
		// 登録されていたら既存のものを利用
		h = found
	} else {
		// 登録されていなかったら新しく用意する
		h = hub.NewHub(roomId)
		(*r.Hubs)[roomId] = h
		go h.Run()
	}
	return h
}

func (r *FocusRepository) CheckHubExists(roomId entity.RoomId) error {
	if _, ok := (*r.Hubs)[roomId]; !ok {
		// 登録されていなかったら
		return fmt.Errorf("FocusRepository.CheckHubExists Error : room not found")
	}
	// 登録されていたら部屋を削除
	delete(*r.Hubs,roomId)
	return nil
}
