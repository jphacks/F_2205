package entity

type Name string

type Member struct {
	Name     Name `json:"name"`
	Connects []*Connect
}

type Members []*Member

type Connect struct {
	Name Name `json:"name"`
}
