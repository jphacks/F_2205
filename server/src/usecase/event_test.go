package usecase

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestEventUsecase_SwitchExecEventByEventType(t *testing.T) {
	// TODO
	t.Skip()
}

func TestEventUsecase_SwitchExecFocusEventByEventType(t *testing.T) {
	// TODO
	t.Skip()
}

func TestEventUsecase_AddNewMemberOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		focusInfo entity.FocusInfo
		wantRoom  *entity.Room
		wantErr   error
	}{
		{
			name:   "正常に動いている場合",
			roomId: entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{
				From: entity.Name("hoge"),
			},
			wantErr: nil,
		},
		{
			name:      "Fromが空文字の場合、エラーを返す",
			roomId:    entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{},
			wantErr:   fmt.Errorf("EventUsecase.AddNewMemberOfRoomId Error : from is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			err := ucEvent.AddNewMemberOfRoomId(tt.roomId, tt.focusInfo)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestEventUsecase_AddNewMemberOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

func TestEventUsecase_SetMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		focusInfo entity.FocusInfo
		wantRoom  *entity.Room
		wantErr   error
	}{
		{
			name:   "正常に動いている場合",
			roomId: entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{
				From: entity.Name("hoge"),
				To:   entity.Name("hoge"),
			},
			wantErr: nil,
		},
		{
			name:   "Fromが空文字の場合、エラーを返す",
			roomId: entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{
				To: entity.Name("hoge"),
			},
			wantErr: fmt.Errorf("EventUsecase.SetMemberFocusOfRoomId Error : from and to is required"),
		},
		{
			name:   "Toが空文字の場合、エラーを返す",
			roomId: entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{
				From: entity.Name("hoge"),
			},
			wantErr: fmt.Errorf("EventUsecase.SetMemberFocusOfRoomId Error : from and to is required"),
		},
		{
			name:      "両方空文字の場合、エラーを返す",
			roomId:    entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{},
			wantErr:   fmt.Errorf("EventUsecase.SetMemberFocusOfRoomId Error : from and to is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			err := ucEvent.SetMemberFocusOfRoomId(tt.roomId, tt.focusInfo)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestEventUsecase_SetMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

func TestEventUsecase_DelMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		focusInfo entity.FocusInfo
		wantRoom  *entity.Room
		wantErr   error
	}{
		{
			name:   "正常に動いている場合",
			roomId: entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{
				From: entity.Name("hoge"),
				To:   entity.Name("hoge"),
			},
			wantErr: nil,
		},
		{
			name:   "Fromが空文字の場合、エラーを返す",
			roomId: entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{
				To: entity.Name("hoge"),
			},
			wantErr: fmt.Errorf("EventUsecase.DelMemberFocusOfRoomId Error : from and to is required"),
		},
		{
			name:   "Toが空文字の場合、エラーを返す",
			roomId: entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{
				From: entity.Name("hoge"),
			},
			wantErr: fmt.Errorf("EventUsecase.DelMemberFocusOfRoomId Error : from and to is required"),
		},
		{
			name:      "両方空文字の場合、エラーを返す",
			roomId:    entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{},
			wantErr:   fmt.Errorf("EventUsecase.DelMemberFocusOfRoomId Error : from and to is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			err := ucEvent.DelMemberFocusOfRoomId(tt.roomId, tt.focusInfo)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestEventUsecase_DelMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

func TestEventUsecase_DelAllMemberFocusOfRoomId(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		roomId    entity.RoomId
		focusInfo entity.FocusInfo
		wantRoom  *entity.Room
		wantErr   error
	}{
		{
			name:   "正常に動いている場合",
			roomId: entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{
				From: entity.Name("hoge"),
			},
			wantErr: nil,
		},
		{
			name:      "Fromが空文字の場合、エラーを返す",
			roomId:    entity.RoomId("1234"),
			focusInfo: entity.FocusInfo{},
			wantErr:   fmt.Errorf("EventUsecase.DelAllMemberFocusOfRoomId Error : from is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			err := ucEvent.DelAllMemberFocusOfRoomId(tt.roomId, tt.focusInfo)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestEventUsecase_DelAllMemberFocusOfRoomId Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

func TestEventUsecase_ExecEffectEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		roomId           entity.RoomId
		effectInfo       entity.EffectInfo
		wantEffectMember *entity.EffectMember
		wantErr          error
	}{
		{
			name:   "正常に動いている場合",
			roomId: entity.RoomId("1234"),
			effectInfo: entity.EffectInfo{
				Name: entity.Name("hoge"),
				Type: entity.EffectType("happy"),
			},
			wantEffectMember: &entity.EffectMember{
				Name: entity.Name("hoge"),
				Type: entity.EffectType("happy"),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucEvent := NewEventUsecase(&roomRepositoryMock{})
			got, err := ucEvent.ExecEffectEvent(tt.roomId, tt.effectInfo)
			if !reflect.DeepEqual(tt.wantEffectMember, got) {
				t.Errorf("TestEventUsecase_ExecEffectEvent Error : want %v, but got %v", tt.wantEffectMember, got)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestEventUsecase_ExecEffectEvent Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}
