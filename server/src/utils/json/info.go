package json

import "github.com/jphacks/F_2205/server/src/domain/entity"

type InfoJson struct {
	Focus  FocusInfoJson    `json:"focus"`
	Effect EffectInfoJson   `json:"effect"`
	Sound  SoundInfoJson    `json:"sound"`
	Member MemberInfoJson   `json:"member"`
	Rest   RestRoomInfoJson `json:"rest"`
}

type FocusInfoJson struct {
	From PeerIdJson `json:"from"`
	To   PeerIdJson `json:"to"`
}

type EffectInfoJson struct {
	PeerId PeerIdJson     `json:"peer_id"`
	Type   EffectTypeJson `json:"type"`
}

type SoundInfoJson struct {
	SoundType SoundTypeJson `json:"type"`
}

type MemberInfoJson struct {
	PeerId PeerIdJson `json:"peer_id"`
	Name   NameJson   `json:"name"`
}

type RestRoomInfoJson struct {
	PeerId     PeerIdJson `json:"peer_id"`
	IsRestRoom bool       `json:"is_rest_room"`
}

type RoomInfoJson struct {
	Id RoomIdJson `json:"id"`
}

// json to entity
func InfoJsonToEntity(j InfoJson) entity.Info {
	return entity.Info{
		Focus:  FocusInfoJsonToEntity(j.Focus),
		Effect: EffectInfoJsonToEntity(j.Effect),
		Sound:  SoundInfoJsonToEntity(j.Sound),
		Member: MemberInfoJsonToEntity(j.Member),
		Rest:   RestRoomInfoJsonToEntity(j.Rest),
	}
}

func FocusInfoJsonToEntity(j FocusInfoJson) entity.FocusInfo {
	return entity.FocusInfo{
		From: entity.PeerId(j.From),
		To:   entity.PeerId(j.To),
	}
}

func EffectInfoJsonToEntity(j EffectInfoJson) entity.EffectInfo {
	return entity.EffectInfo{
		PeerId: entity.PeerId(j.PeerId),
		Type:   entity.EffectType(j.Type),
	}
}

func SoundInfoJsonToEntity(j SoundInfoJson) entity.SoundInfo {
	return entity.SoundInfo{
		SoundType: entity.SoundType(j.SoundType),
	}
}

func MemberInfoJsonToEntity(j MemberInfoJson) entity.MemberInfo {
	return entity.MemberInfo{
		PeerId: entity.PeerId(j.PeerId),
		Name:   entity.Name(j.Name),
	}
}

func RestRoomInfoJsonToEntity(j RestRoomInfoJson) entity.RestRoomInfo {
	return entity.RestRoomInfo{
		PeerId:     entity.PeerId(j.PeerId),
		IsRestRoom: j.IsRestRoom,
	}
}

func RoomInfoJsonToEntity(j RoomInfoJson) entity.RoomInfo {
	return entity.RoomInfo{
		Id: entity.RoomId(j.Id),
	}
}

// entity to json
func RoomInfoEntityToJson(r *entity.RoomInfo) RoomInfoJson {
	return RoomInfoJson{
		Id: RoomIdJson(r.Id),
	}
}
