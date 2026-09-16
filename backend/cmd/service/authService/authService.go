package authService

import (
	"errors"
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	storage    UserStorage
	jwtService JWTService
}
type JWTService interface {
	GenerateToken(userID string) (string, error)
}
type UserStorage interface {
	FindUser(email string) (*pm_entity.User, error)
}

func NewAuthService(storage UserStorage, jwtService JWTService) *AuthService {
	return &AuthService{
		storage:    storage,
		jwtService: jwtService,
	}
}
func (as *AuthService) Login(model userDTO.LoginDTO) (*string, *app_errors.HttpError) {
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

	return &token, nil
}
