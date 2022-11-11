package json

import "github.com/jphacks/F_2205/server/src/domain/entity"

type NameJson string

type PeerIdJson string

type MemberJson struct {
	Name       NameJson `json:"name"`
	IsRestRoom bool     `json:"isRestRoom"`
}

type MembersJson map[PeerIdJson]*MemberJson

// entity to json
func MemberEntityToJson(m *entity.Member) *MemberJson {
	return &MemberJson{
		Name:       NameJson(m.Name),
		IsRestRoom: m.IsRestRoom,
	}
}

func MembersEntityToJson(ms entity.Members) MembersJson {
	r := MembersJson{}
	for key, value := range ms {
		r[PeerIdJson(key)] = MemberEntityToJson(value)
	}
	return r
}
