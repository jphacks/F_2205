package json

import "github.com/jphacks/F_2205/server/src/domain/entity"

type EventTypeJson string

type EventJson struct {
	Type EventTypeJson `json:"type"`
	Info InfoJson      `json:"info"`
}

// json to entity
func EventJsonToEntity(e EventJson) entity.Event {
	return entity.Event{
		Type: entity.EventType(e.Type),
		Info: InfoJsonToEntity(e.Info),
	}
}
