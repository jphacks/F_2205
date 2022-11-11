package json

import "github.com/jphacks/F_2205/server/src/domain/entity"

type EffectTypeJson string

type EffectMemberJson struct {
	PeerId PeerIdJson     `json:"peer_id"`
	Type   EffectTypeJson `json:"type"`
}

// entity to json
func EffectMemberEntityToJson(e entity.EffectMember) EffectMemberJson {
	return EffectMemberJson{
		PeerId: PeerIdJson(e.PeerId),
		Type:   EffectTypeJson(e.Type),
	}
}
