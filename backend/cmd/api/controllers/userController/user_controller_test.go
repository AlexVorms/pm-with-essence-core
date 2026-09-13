package userController

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pm-with-essence/cmd/api/model/responseDTO"
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/internal/api/app_errors"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserController_Register_Success(t *testing.T) {
	//ARRANGE
	var receivedDTO userDTO.RegisterDTO
	mockService := MockUserService{
		CreateNewUserFunc: func(dto userDTO.RegisterDTO) *app_errors.HttpError {
			receivedDTO = dto
			return nil
		},
	}
	newUserDto := userDTO.RegisterDTO{
		UserName: "testUser",
		Password: "Password123",
		Email:    "user1112@mail.ru",
	}
	body, err := json.Marshal(newUserDto)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer(body),
	)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	//ACT
	controller := NewUserController(&mockService)
	controller.Register(c)

	//ASSERT
	if w.Code != http.StatusOK {
		t.Errorf("Wanted status code %v, got %v", http.StatusOK, w.Code)
	}

	var receivedResponse responseDTO.MessageResponse
	err = json.Unmarshal(w.Body.Bytes(), &receivedResponse)
	if err != nil {
		t.Fatalf("Ошибка десериализации JSON: %v", err)
	}
	expectedResponse := responseDTO.MessageResponse{Message: "success"}
	if receivedResponse != expectedResponse {
		t.Errorf("Wanted body %v, got %v", expectedResponse, receivedResponse)
	}

	if !mockService.CreateNewUserCalled {
		t.Errorf("CreateNewUser was not called")
	}
	if receivedDTO.UserName != newUserDto.UserName {
		t.Errorf("Wanted body %v, got %v", newUserDto.UserName, receivedDTO.UserName)
	}
	if receivedDTO.Password != newUserDto.Password {
		t.Errorf("Wanted body %v, got %v", newUserDto.Password, receivedDTO.Password)
	}
	if receivedDTO.Email != newUserDto.Email {
		t.Errorf("Wanted body %v, got %v", newUserDto.Email, receivedDTO.Email)
	}
}
