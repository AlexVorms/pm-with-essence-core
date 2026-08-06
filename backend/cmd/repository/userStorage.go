package repository

import (
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"gorm.io/gorm"
)

type UserStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) FindUser(email string) (*pm_entity.User, error) {
	var user pm_entity.User

	err := s.db.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (s *UserStorage) CreateNewUser(user pm_entity.User) error {
	return s.db.Model(pm_entity.User{}).Create(&user).Error
}
