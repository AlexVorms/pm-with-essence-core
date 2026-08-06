package authService

import (
	"errors"
	"pm-with-essence/cmd/api/model/responseDTO"
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/internal/api/app_errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	storage    *repository.UserStorage
	jwtService *JwtService
}

func NewAuthService(storage *repository.UserStorage, jwtService *JwtService) *AuthService {
	return &AuthService{
		storage:    storage,
		jwtService: jwtService,
	}
}
func (as *AuthService) Login(model userDTO.LoginDTO) (*responseDTO.LoginResponse, *app_errors.HttpError) {
	user, err := as.storage.FindUser(model.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app_errors.NewLoginError()
	}

	if err != nil {
		return nil, app_errors.NewPostgresReadError(
			err,
			"Failed to read user from database",
		)
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.HashPassword),
		[]byte(model.Password))

	if err != nil {
		return nil, app_errors.NewLoginError()
	}

	token, err := as.jwtService.GenerateToken(user.ID.String())
	if err != nil {
		return nil, nil
	}

	return &responseDTO.LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}, nil
}
