package userService

import (
	"errors"
	"net/http"
	"pm-with-essence/cmd/api/model/userDTO"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestUserService_IsUserExist_UserExist(t *testing.T) {
	mockStorage := &MockUserStorage{
		FindUserFunc: func(email string) (*pm_entity.User, error) {
			return &pm_entity.User{}, nil
		},
	}
	service := NewUserService(mockStorage)
	err := service.IsUserExist("user1@mail.ru")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", err.Code)
	}
}
func TestUserService_IsUserExist_UserNotFound(t *testing.T) {
	mockStorage := &MockUserStorage{
		FindUserFunc: func(email string) (*pm_entity.User, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}
	service := NewUserService(mockStorage)
	err := service.IsUserExist("user1@mail.ru")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

}
func TestUserService_IsUserExist_StorageError(t *testing.T) {
	mockStorage := &MockUserStorage{
		FindUserFunc: func(email string) (*pm_entity.User, error) {
			return nil, errors.New("database connection failed")
		},
	}
	service := NewUserService(mockStorage)
	err := service.IsUserExist("user1@mail.ru")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", err.Code)
	}
}
func TestUserService_CreateNewUser_UserExist(t *testing.T) {
	mockStorage := &MockUserStorage{
		FindUserFunc: func(email string) (*pm_entity.User, error) {
			return &pm_entity.User{}, nil
		},
	}
	service := NewUserService(mockStorage)
	err := service.CreateNewUser(userDTO.RegisterDTO{
		UserName: "user1",
		Password: "Password123",
		Email:    "user111@mail.ru",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", err.Code)
	}
}
func TestUserService_CreateNewUser_InvalidPassword(t *testing.T) {
	mockStorage := &MockUserStorage{
		FindUserFunc: func(email string) (*pm_entity.User, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}
	service := NewUserService(mockStorage)
	err := service.CreateNewUser(userDTO.RegisterDTO{
		UserName: "user1",
		Password: "password123",
		Email:    "User111@mail.ru",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", err.Code)
	}
	if err.Message != "Password must contain uppercase, lowercase and digit" {
		t.Errorf("unexpected error message: %s", err.Message)
	}
}
func TestUserService_CreateNewUser_Success(t *testing.T) {
	var createdUser *pm_entity.User

	mockStorage := &MockUserStorage{
		FindUserFunc: func(email string) (*pm_entity.User, error) {
			return nil, gorm.ErrRecordNotFound
		},
		CreateNewUserFunc: func(user *pm_entity.User) error {
			createdUser = user
			return nil
		},
	}
	service := NewUserService(mockStorage)
	user := userDTO.RegisterDTO{
		UserName: "user1",
		Password: "Password123",
		Email:    "User111@mail.ru",
	}
	expectedEmail := "user111@mail.ru"

	err := service.CreateNewUser(user)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	if createdUser == nil {
		t.Fatal("expected user, got nil")
	}
	if createdUser.Name != user.UserName {
		t.Errorf("expected username %s, got %s", user.UserName, createdUser.Name)
	}
	if createdUser.Email != expectedEmail {
		t.Errorf("expected email %s, got %s", user.Email, createdUser.Email)
	}
	if createdUser.HashPassword == user.Password {
		t.Fatal("expected hashed password to be different")
	}
	errPassword := bcrypt.CompareHashAndPassword(
		[]byte(createdUser.HashPassword),
		[]byte(user.Password),
	)

	if errPassword != nil {
		t.Errorf("password hash does not match original password")
	}
}
func TestUserService_CreateNewUser_StorageError(t *testing.T) {
	mockStorage := &MockUserStorage{
		FindUserFunc: func(email string) (*pm_entity.User, error) {
			return nil, gorm.ErrRecordNotFound
		},
		CreateNewUserFunc: func(user *pm_entity.User) error {
			return errors.New("database write error")
		},
	}
	service := NewUserService(mockStorage)
	user := userDTO.RegisterDTO{
		UserName: "user1",
		Password: "Password123",
		Email:    "User111@mail.ru",
	}
	err := service.CreateNewUser(user)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", err.Code)
	}
}
