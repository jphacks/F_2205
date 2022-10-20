package entity

// RoomはTodoリストのタスクに関する情報を保持する構造体です
type Room struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// RoomsはRoomのリストの型です
type Rooms []*Room

type RoomId string
