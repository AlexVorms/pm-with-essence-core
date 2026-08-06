package pm_entity

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name         string    `json:"name" gorm:"not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	HashPassword string    `json:"password" gorm:"not null"`
}

func CreateUserEntity(name string, password string, email string) User {
	return User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		HashPassword: password,
	}
}
