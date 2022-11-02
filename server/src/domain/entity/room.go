package entity

type RoomId string

type Room struct {
	EventType  EventType  `json:"event_type"`
	Focus      Focus      `json:"focus"`
}

type Rooms map[RoomId]*Room
