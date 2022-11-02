package entity

// EffectTypeはMemberのエフェクトのタイプです
type EffectType string

const (
	Drink EffectType = "DRINK"
)

// EffectMemberはどのMemberがどのエフェクトを使ったかの情報を保持します
type EffectMember struct {
	Name Name       `json:"name"`
	Type EffectType `json:"type"`
}
