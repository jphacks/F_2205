package entity

import "sync"

// FocusMemberはNameというMemberがどのMemberとつながっているかConnectsに保持しています
type FocusMember struct {
	PeerId   PeerId
	Connects Connects
	Mu       sync.RWMutex // Connectsをロックする
}

type FocusMembers []*FocusMember

type FocusMembersStore struct {
	FocusMembers FocusMembers
	Mu           sync.RWMutex // FocusMembersをロックする
}

// ConnectはどのMemberとFocus状態にあるかを管理します
type Connect struct {
	PeerId PeerId
}

type Connects []*Connect
