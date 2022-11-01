package entity

type EventType string

type Event struct {
	Type  EventType `json:"type"`
	Info  Info      `json:"info"`
	Focus Focus     `json:"focus"`
}

type Events map[RoomId]*Event

const (
	NewMember   EventType = "NEW_MEMBER"    //fromを新しいmemberとして追加する
	SetFocus    EventType = "SET_FOCUS"     //from と toをつなげる
	DelFocus    EventType = "DEL_FOCUS"     //from と toを解除する
	DelAllFocus EventType = "DEL_ALL_FOCUS" //fromのすべてのconnectを削除する
)
