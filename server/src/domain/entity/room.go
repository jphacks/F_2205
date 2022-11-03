package entity

type RoomId string

type Room struct {
	EventType    EventType    `json:"event_type"`
	FocusMembers FocusMembers `json:"focus_members"`
	EffectMember EffectMember `json:"effect_member"`
}

type Rooms map[RoomId]*Room

type RoomInfo struct {
	Id RoomId `json:"id"`
}
