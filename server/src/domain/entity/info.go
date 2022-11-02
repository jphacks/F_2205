package entity

// InfoはEventの情報を管理する構造体です
type Info struct {
	Focus  FocusInfo  `json:"focus"`
	Effect EffectInfo `json:"effect"`
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
