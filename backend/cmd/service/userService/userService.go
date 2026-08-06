package userService

import (
	"errors"
	"fmt"
	"pm-with-essence/cmd/api/model/userDTO"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"
	"pm-with-essence/internal/validator"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	storage *repository.UserStorage
}

func NewUserService(storage *repository.UserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}
func (us *UserService) CreateNewUser(dto userDTO.RegisterDTO) *app_errors.HttpError {
	fmt.Printf(dto.Email, dto.Password, dto.UserName)

	if err := validator.ValidatePassword(dto.Password); err != nil {
		return err
	}

	email := strings.TrimSpace(strings.ToLower(dto.Email))

	user, err := us.storage.FindUser(email)
	fmt.Println(user)

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

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(dto.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return app_errors.NewPasswordHashError(err, "Failed to hash password")
	}
	println(string(hashPassword))
	newUser := pm_entity.CreateUserEntity(dto.UserName, string(hashPassword), email)
	err = us.storage.CreateNewUser(newUser)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Postgres write user data error")
	}
	return nil
}
