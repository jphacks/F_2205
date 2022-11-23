package entity

// Infoは受け取ったEventの情報を管理する構造体です
type Info struct {
	Focus  FocusInfo
	Effect EffectInfo
	Sound  SoundInfo
	Member MemberInfo
	Rest   RestRoomInfo
}

// FocusInfoはFocusEventの際のMember同士の情報を保持する構造体です
type FocusInfo struct {
	From PeerId
	To   PeerId
}

// EffectInfoはMemberのエフェクトを保持する構造体です
type EffectInfo struct {
	PeerId PeerId
	Type   EffectType
}

// SoundInfoは受け取ったSoundの情報を保持する構造体です
type SoundInfo struct {
	SoundType SoundType
}

// MemberInfoはMemberの情報を保持する構造体です
type MemberInfo struct {
	PeerId PeerId
	Name   Name
}

// RestRoomInfoはお手洗いに行っているかどうかを保持する構造体です
type RestRoomInfo struct {
	PeerId     PeerId
	IsRestRoom bool
}

// RoomInfoはRoomのメタデータを保持します
type RoomInfo struct {
	Id RoomId
}
