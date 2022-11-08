package entity

// Infoは受け取ったEventの情報を管理する構造体です
type Info struct {
	Focus  FocusInfo    `json:"focus"`
	Effect EffectInfo   `json:"effect"`
	Sound  SoundInfo    `json:"sound"`
	Member MemberInfo   `json:"member"`
	Rest   RestRoomInfo `json:"rest"`
}

// FocusInfoはFocusEventの際のMember同士の情報を保持する構造体です
type FocusInfo struct {
	From   PeerId `json:"from"`
	To     PeerId `json:"to"`
}

// EffectInfoはMemberのエフェクトを保持する構造体です
type EffectInfo struct {
	PeerId PeerId   `json:"peer_id"`
	Type EffectType `json:"type"`
}

// SoundInfoは受け取ったSoundの情報を保持する構造体です
type SoundInfo struct {
	SoundType SoundType `json:"type"`
}

// MemberInfoはMemberの情報を保持する構造体です
type MemberInfo struct {
	PeerId PeerId `json:"peer_id"`
	Name   Name   `json:"name"`
}

// RestRoomInfoはお手洗いに行っているかどうかを保持する構造体です
type RestRoomInfo struct {
	PeerId     PeerId `json:"peer_id"`
	IsRestRoom bool   `json:"is_rest_room"`
}

// RoomInfoはRoomのメタデータを保持します
type RoomInfo struct {
	Id RoomId
}
