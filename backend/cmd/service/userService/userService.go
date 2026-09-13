package userService

import (
	"errors"
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"
	"pm-with-essence/internal/validator"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	storage UserStorage
}
type UserStorage interface {
	FindUser(email string) (*pm_entity.User, error)
	CreateNewUser(user *pm_entity.User) error
}

func NewUserService(storage UserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}
func (us *UserService) CreateNewUser(dto userDTO.RegisterDTO) *app_errors.HttpError {
	email := strings.TrimSpace(strings.ToLower(dto.Email))

	err := us.IsUserExist(email)
	if err != nil {
		return err
	}

	if err = validator.ValidatePassword(dto.Password); err != nil {
		return err
	}

	hashPassword, err1 := bcrypt.GenerateFromPassword(
		[]byte(dto.Password),
		bcrypt.DefaultCost,
	)
	if err1 != nil {
		return app_errors.NewPasswordHashError(err1, "Failed to hash password")
	}

	newUser := pm_entity.CreateUserEntity(dto.UserName, string(hashPassword), email)

	err2 := us.storage.CreateNewUser(&newUser)
	if err2 != nil {
		return app_errors.NewPostgresWriteError(err2, "Postgres write user data error")
	}

	return nil
}

func (us *UserService) IsUserExist(email string) *app_errors.HttpError {
	_, err := us.storage.FindUser(email)

	if err == nil {
		return app_errors.NewBadRequestError(
			"User with this email already exists",
		)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return app_errors.NewPostgresReadError(
			err,
			"Failed to read user from database",
		)
	}
	return nil
}
