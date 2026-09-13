package userController

import (
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/internal/api/app_errors"
)

type MockUserService struct {
	CreateNewUserFunc   func(dto userDTO.RegisterDTO) *app_errors.HttpError
	CreateNewUserCalled bool
}

func (m *MockUserService) CreateNewUser(dto userDTO.RegisterDTO) *app_errors.HttpError {
	m.CreateNewUserCalled = true
	return m.CreateNewUserFunc(dto)
}
