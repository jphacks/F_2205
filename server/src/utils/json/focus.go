package json

import (
	"github.com/jphacks/F_2205/server/src/domain/entity"
)

type FocusMemberJson struct {
	PeerId   entity.PeerId `json:"peer_id"`
	Connects ConnectsJson  `json:"connects"`
}

type FocusMembersJson []*FocusMemberJson

type ConnectJson struct {
	PeerId entity.PeerId `json:"peer_id"`
}

type ConnectsJson []*ConnectJson

// entity to json
func FocusMembersEntityToJson(ms entity.FocusMembers) FocusMembersJson {
	r := FocusMembersJson{}
	for _, m := range ms {
		r = append(r, FocusMemberEntityToJson(m))
	}
	return r
}

func FocusMemberEntityToJson(m *entity.FocusMember) *FocusMemberJson {
	m.Mu.RLock()
	defer m.Mu.RUnlock()
	return &FocusMemberJson{
		PeerId:   m.PeerId,
		Connects: ConnectsEntityToJson(m.Connects),
	}
}

func ConnectsEntityToJson(cs entity.Connects) ConnectsJson {
	r := ConnectsJson{}
	for _, c := range cs {
		r = append(r, ConnectEntityToJson(c))
	}
	return r
}

func ConnectEntityToJson(c *entity.Connect) *ConnectJson {
	return &ConnectJson{
		PeerId: c.PeerId,
	}
}
