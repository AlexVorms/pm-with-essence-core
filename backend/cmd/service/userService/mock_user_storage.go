package userService

import pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

type MockUserStorage struct {
	FindUserFunc      func(email string) (*pm_entity.User, error)
	CreateNewUserFunc func(user *pm_entity.User) error
}

func (m *MockUserStorage) FindUser(email string) (*pm_entity.User, error) {
	return m.FindUserFunc(email)
}
func (m *MockUserStorage) CreateNewUser(user *pm_entity.User) error {
	return m.CreateNewUserFunc(user)
}
