package usecase

import (
	"fmt"

	"github.com/jphacks/F_2205/server/src/domain/entity"
	"github.com/jphacks/F_2205/server/src/domain/repository"
	"github.com/jphacks/F_2205/server/src/utils/generate"
)

// RoomUsecaseはユースケースの構造体です
type RoomUsecase struct {
	repo repository.IRoomRepository
}

// IRoomUsecaseはRoomのユースケースをまとめたインターフェースです
type IRoomUsecase interface {
	GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers
	GetMembersOfRoomId(roomId entity.RoomId) entity.Members
	CheckExistsRoomAndInit(roomId entity.RoomId)
	DeleteRoomOfRoomId(roomId entity.RoomId)
	CreateRoomNumber() (*entity.RoomInfo, error)
	GetSumOfRoom() int
	SetRoomLatestMemberDataOfRoomId(roomID entity.RoomId,room *entity.Room, e entity.Event)
}

// NewRoomUsecaseはIRoomUsecaseを満たしたRoomUsecase構造体を返します
func NewRoomUsecase(repo repository.IRoomRepository) IRoomUsecase {
	return &RoomUsecase{
		repo: repo,
	}
}

// CreateRoomNumberはRoom番号を生成します
// またそのRoom番号がすでに使われているか確認し、使われていた場合エラーを返します
func (uc *RoomUsecase) CreateRoomNumber() (*entity.RoomInfo, error) {
	roomIdString, err := generate.MakeRandomStrFromLetters(4)
	if err != nil {
		return nil, fmt.Errorf("RoomUsecase.CreateRoomNumber Error : %w", err)
	}
	roomId := (entity.RoomId)(roomIdString)
	_, found := uc.repo.GetExistsRoomOfRoomId(roomId)
	if found {
		return nil, fmt.Errorf("RoomUsecase.CreateRoomNumber Error : roomId already used")
	}
	roomInfo := &entity.RoomInfo{
		Id: roomId,
	}
	return roomInfo, nil
}

// DeleteRoomOfRoomIdは受け取ったIDのRoomを削除します
func (uc *RoomUsecase) DeleteRoomOfRoomId(roomId entity.RoomId) {
	uc.repo.DeleteRoomOfRoomId(roomId)
}

// GetFocusMembersOfRoomIdは受け取ったIDのRoomのフォーカスメンバーを返します
func (uc *RoomUsecase) GetFocusMembersOfRoomId(roomId entity.RoomId) entity.FocusMembers {
	return uc.repo.GetFocusMembersOfRoomId(roomId)
}

// GetMembersOfRoomIdは受け取ったIDのRoomのメンバーを返します
func (uc *RoomUsecase) GetMembersOfRoomId(roomId entity.RoomId) entity.Members {
	return uc.repo.GetMembersOfRoomId(roomId)
}

// CheckExistsRoomAndInitは受け取ったIDのRoomが登録されているか確認し、
// 登録されてなかった場合Roomを初期化して用意します
func (uc *RoomUsecase) CheckExistsRoomAndInit(roomId entity.RoomId) {
	uc.repo.CheckExistsRoomAndInit(roomId)
}

// GetSumOfRoomは存在するRoomの合計数を返します
func (uc *RoomUsecase) GetSumOfRoom() int {
	return uc.repo.GetSumOfRoom()
}

func (uc *RoomUsecase) SetRoomLatestMemberDataOfRoomId(roomID entity.RoomId,room *entity.Room, e entity.Event) {
	room.FocusMembers = uc.GetFocusMembersOfRoomId(roomID)
	room.Members = uc.GetMembersOfRoomId(roomID)
}