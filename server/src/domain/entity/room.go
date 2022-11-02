package entity

type RoomId string

type Room struct {
	Focus Focus `json:"focus"`
}

type Rooms map[RoomId]*Room
