package entity

type Name string

// MemberはRoomにいるユーザーの情報を保持します
type Member struct {
	PeerId string `json:"peer_id"`
	Name   Name   `json:"name"`
}

type Members []*Member
