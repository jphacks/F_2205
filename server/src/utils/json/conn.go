package json

type CountRoomConnJson struct {
	Count int `json:"count"`
}

func CreatecountRoomConnJson(count int) CountRoomConnJson {
	return CountRoomConnJson{
		Count: count,
	}
}
