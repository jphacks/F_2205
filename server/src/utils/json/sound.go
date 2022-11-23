package json

import "github.com/jphacks/F_2205/server/src/domain/entity"

type SoundTypeJson string

type SoundJson struct {
	SoundType SoundTypeJson `json:"type"`
}

// entity to json
func SoundEntityToJson(s entity.Sound) SoundJson {
	return SoundJson{
		SoundType: SoundTypeJson(s.SoundType),
	}
}
