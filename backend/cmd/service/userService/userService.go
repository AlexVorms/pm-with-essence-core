package userService

import "pm-with-essence/cmd/repository"

type UserService struct {
	storage *repository.UserStorage
}

func NewUserService(storage *repository.UserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}
