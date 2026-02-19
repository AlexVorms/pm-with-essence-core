package service

import (
	"pm-with-essence/internal/api/errors"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (service *Service) Run() *errors.HttpError {
	var err error
	return errors.NewCreateJWTError(err)
}
