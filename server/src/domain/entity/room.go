package entity

// RoomはTodoリストのタスクに関する情報を保持する構造体です
type Room struct {
	Id   int
	Name string
}

// RoomsはRoomのリストの型です
type Rooms []*Room
