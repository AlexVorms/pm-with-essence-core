package service

import (
	"pm-with-essence/internal/api/errors"

	"gorm.io/gorm"
)

type Service struct {
	*gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{}
}
func (service *Service) Run() *errors.HttpError {
	var err error
	return errors.NewCreateJWTError(err)
}
