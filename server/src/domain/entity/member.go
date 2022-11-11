package entity

import "sync"

type Name string

type PeerId string

// MemberはRoomにいるユーザーの情報を保持します
type Member struct {
	Name       Name
	IsRestRoom bool
}

type Members map[PeerId]*Member

type MembersStore struct {
	Members Members
	Mu      sync.RWMutex
}
