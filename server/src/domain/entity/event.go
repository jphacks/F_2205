package entity

// EventTypeはwebsocket通信の際のEventを管理する構造体です
type EventType string

// TODO ドメインはjsonに依存すべきではない
type Event struct {
	Type EventType `json:"type"`
	Info Info      `json:"info"`
}

type Events map[RoomId]*Event

const (
	// FocusEvent
	NewMember     EventType = "NEW_MEMBER"      // fromを新しいmemberとして追加する
	SetFocus      EventType = "SET_FOCUS"       // from と toをつなげる
	DelFocus      EventType = "DEL_FOCUS"       // from と toを解除する
	DelAllFocus   EventType = "DEL_ALL_FOCUS"   // fromのすべてのconnectを削除する

	// ScreenShotEvent
	SetScreenShot EventType = "SET_SCREEN_SHOT" // スクリーンショットの開始

	// EffectEvent
	SetEffect     EventType = "SET_EFFECT"      // エフェクトを設定する

	// SoundEvent
	SetSound      EventType = "SET_SOUND"       // サウンドの設定をする
)
