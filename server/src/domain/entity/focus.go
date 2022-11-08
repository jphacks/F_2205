package entity

// FocusMemberはNameというMemberがどのMemberとつながっているかConnectsに保持しています
type FocusMember struct {
	PeerId   PeerId   `json:"peer_id"`
	Connects Connects `json:"connects"`
}

type FocusMembers []*FocusMember

// ConnectはどのMemberとFocus状態にあるかを管理します
type Connect struct {
	PeerId PeerId `json:"peer_id"`
}

type Connects []*Connect
