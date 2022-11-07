package entity

type Name string

type PeerId string

// MemberはRoomにいるユーザーの情報を保持します
type Member struct {
	Name         Name    `json:"name"`
	IsRestRoom   bool    `json:"isRestRoom"`
}

type Members map[PeerId]*Member
