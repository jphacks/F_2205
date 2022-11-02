package entity

// FocusMemberはNameというMemberがどのMemberとつながっているかConnectsに保持しています
type FocusMember struct {
	Name     Name     `json:"name"`
	Connects Connects `json:"connects"`
}

type FocusMembers []*FocusMember

// ConnectはどのMemberとFocus状態にあるかを管理します
type Connect struct {
	Name Name `json:"name"`
}

type Connects []*Connect
