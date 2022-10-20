package entity

type EventType string

type Event struct {
	Type EventType `json:"type"`
	Info Info `json:"info"`
	Focus Focus `json:"focus"`
}

const (
	NewMember EventType = "NEW_MEMBER" //新しいmemberをmembers配列に追加する
	SetFocus EventType = "SET_FOCUS" //Focusをセットする
)
