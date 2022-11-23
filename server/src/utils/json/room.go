package json

import "github.com/jphacks/F_2205/server/src/domain/entity"

type RoomIdJson string

// Roomは部屋で起こったイベントの情報やステートを保持します
type RoomJson struct {
	EventType    EventTypeJson    `json:"event_type"`
	Members      MembersJson      `json:"members"`
	FocusMembers FocusMembersJson `json:"focus_members"`
	EffectMember EffectMemberJson `json:"effect_member"`
	Sound        SoundJson        `json:"sound"`
}

type CountRoomJson struct {
	Count int `json:"count"`
}

// entity to json
func RoomEntityToJson(r *entity.Room) *RoomJson {
	r.MembersStore.Mu.RLock()
	defer r.MembersStore.Mu.RUnlock()

	r.FocusMembersStore.Mu.RLock()
	defer r.FocusMembersStore.Mu.RUnlock()

	return &RoomJson{
		EventType:    EventTypeJson(r.EventType),
		Members:      MembersEntityToJson(r.MembersStore.Members),
		FocusMembers: FocusMembersEntityToJson(r.FocusMembersStore.FocusMembers),
		EffectMember: EffectMemberEntityToJson(r.EffectMember),
		Sound:        SoundEntityToJson(r.Sound),
	}
}

func RoomIdEntityToJson(r entity.RoomId) RoomIdJson {
	return RoomIdJson(r)
}

// json to entity
func RoomIdJsonToEntity(r RoomIdJson) entity.RoomId {
	return entity.RoomId(r)
}

// create response
func CreateCountRoomJson(count int) CountRoomJson {
	return CountRoomJson{
		Count: count,
	}
}
