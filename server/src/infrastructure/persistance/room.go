package persistance

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
	"github.com/jphacks/F_2205/server/src/infrastructure/database"
)

// RoomRepositoryがdomainのinterfaceを満たしているか確認します
var _ repository.IRoomRepository = &RoomRepository{}

// RoomRepositoryはrepository.RoomRepositoryを満たす構造体です
type RoomRepository struct {
	conn *database.Conn
}

// NewRoomRepositoryはRoomRepositoryのポインタを生成する関数です
func NewRoomRepository(conn *database.Conn) *RoomRepository {
	return &RoomRepository{
		conn: conn,
	}
}

// CreateRoomは受け取ったroomをDBに保存します
func (r *RoomRepository) CreateRoom(room *entity.Room) (*entity.Room, error) {
	query := `INSERT INTO rooms (name) VALUES (:name)`

	dto := roomEntityToDTO(room)
	_, err := r.conn.DB.NamedExec(query, dto)
	if err != nil {
		return nil, fmt.Errorf("RoomRepositoy.CreateRoom NamedExec Error : %w", err)
	}
	return roomDTOtoEntity(dto), nil
}

// entityのroomのデータをやり取りするオブジェクトです
// またDTOとは、Data Transfer Objectの略です
type roomDTO struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

// roomEntityToDTOはroomのentityをdtoに変えます
func roomEntityToDTO(room *entity.Room) roomDTO {
	return roomDTO{
		Id:   room.Id,
		Name: room.Name,
	}
}

// roomEntitiesToDTOsはroomのdtoのリストをentityのリストをに変えます
func roomDTOsToEntities(dtos []roomDTO) entity.Rooms {
	rooms := entity.Rooms{}
	for _, dto := range dtos {
		rooms = append(rooms, &entity.Room{
			Id:   dto.Id,
			Name: dto.Name,
		})
	}
	return rooms
}

// roomDTOtoEntityはroomのdtoをentityに変えます
func roomDTOtoEntity(dto roomDTO) *entity.Room {
	return &entity.Room{
		Id:   dto.Id,
		Name: dto.Name,
	}
}
