package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jphacks/F_2205/server/src/domain/entity"
)

func TestHandler_CreateRoom(t *testing.T) {
	tests := []struct {
		name     string
		reqBody  string
		wantCode int
		wantBody string
	}{
		{
			name:     "正常に動作した場合",
			reqBody:  `{"name":"hoge"}`,
			wantCode: http.StatusOK,
			wantBody: `{"data":{"id":0,"name":"hoge"}}`,
		},
		{
			name:     "request bodyが空だった場合、400エラーになる",
			reqBody:  ``,
			wantCode: http.StatusBadRequest,
			wantBody: `{"error":"RoomHandler.CreateRoom ShouldBindJSON Error : EOF"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()

			roomHandler := NewRoomHandler(&RoomUsecaseMock{})
			router.POST("/room/create", roomHandler.CreateRoom)

			body := bytes.NewBufferString(tt.reqBody)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/room/create", body)
			router.ServeHTTP(w, req)

			if tt.wantCode != w.Code {
				t.Errorf("TestHandler_CreateRoom code Error : want %d but got %d", tt.wantCode, w.Code)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("TestHandler_CreateRoom body Error : want %s but got %s", tt.wantBody, w.Body.String())
			}
		})
	}
}

type RoomUsecaseMock struct{}

func (h *RoomUsecaseMock) CreateRoom(room *entity.Room) (*entity.Room, error) {
	return testCreateRoom, nil
}

// ここでテストデータの用意をします
var (
	testCreateRoom *entity.Room = &entity.Room{
		Id:   0,
		Name: "hoge",
	}
)
