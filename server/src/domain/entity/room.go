package entity

// RoomIdはRoomに割り当てられるプライマリーキーです
type RoomId string

// Roomは部屋で起こったイベントの情報やステートを保持します
type Room struct {
	EventType    EventType    `json:"event_type"`
	Members      Members      `json:"members"`
	FocusMembers FocusMembers `json:"focus_members"`
	EffectMember EffectMember `json:"effect_member"`
}

type Rooms map[RoomId]*Room

// RoomInfoはRoomのメタデータを保持します
type RoomInfo struct {
	Id RoomId
}
