package entity

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name  string    `json:"name" gorm:"unique;not null"`
	Email string    `json:"email" gorm:"unique;not null"`
	//TODO: сделать хэширование паролей
	Password    string `json:"password" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
}
