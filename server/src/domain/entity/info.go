package entity

// Infoは受け取ったEventの情報を管理する構造体です
type Info struct {
	Focus  FocusInfo  `json:"focus"`
	Effect EffectInfo `json:"effect"`
	Sound  SoundInfo  `json:"sound"`
	Member MemberInfo `json:"member"`
}

// FocusInfoはFocusEventの際のMember同士の情報を保持する構造体です
type FocusInfo struct {
	From Name `json:"from"`
	To   Name `json:"to"`
}

// EffectInfoはMemberのエフェクトを保持する構造体です
type EffectInfo struct {
	Name Name       `json:"name"`
	Type EffectType `json:"type"`
}

// SoundInfoは受け取ったSoundの情報を保持する構造体です
type SoundInfo struct {
	SoundType SoundType `json:"type"`
}

type MemberInfo struct {
	PeerId string `json:"peer_id"`
	Name   Name   `json:"name"`
}
