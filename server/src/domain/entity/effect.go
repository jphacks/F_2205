package entity

// EffectTypeはMemberのエフェクトのタイプです
type EffectType string

// EffectMemberはどのMemberがどのエフェクトを使ったかの情報を保持します
type EffectMember struct {
	PeerId PeerId     `json:"peer_id"`
	Type   EffectType `json:"type"`
}
