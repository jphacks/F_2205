package entity

// EventTypeはwebsocket通信の際のEventを管理する構造体です
type EventType string

type Event struct {
	Type EventType
	Info Info
}

const (
	// MembmerEvent
	AddNewMember EventType = "ADD_NEW_MEMBER" // 新しいmemberを部屋に追加する

	// FocusEvent
	SetFocus    EventType = "SET_FOCUS"     // from と toをつなげる
	DelFocus    EventType = "DEL_FOCUS"     // from と toを解除する
	DelAllFocus EventType = "DEL_ALL_FOCUS" // fromのすべてのconnectを削除する

	// ScreenShotEvent
	SetScreenShot EventType = "SET_SCREEN_SHOT" // スクリーンショットの開始

	// EffectEvent
	SetEffect EventType = "SET_EFFECT" // エフェクトを設定する

	// SoundEvent
	SetSound EventType = "SET_SOUND" // サウンドの設定をする

	// RestRoomEvent
	SetRestRoom EventType = "SET_REST_ROOM" // トイレで離席するイベントを設定する
)
