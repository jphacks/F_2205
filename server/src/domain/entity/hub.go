package entity

type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan *Event

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	Focus *Focus

	RoomId RoomId
}

type Hubs map[RoomId]*Hub
