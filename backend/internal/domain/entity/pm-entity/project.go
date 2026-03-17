package pm_entity

import "github.com/google/uuid"

type Project struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name" gorm:"not null"`
	IsPublic    bool      `json:"isPublic"`
	Description string    `json:"description" gorm:"type:text"`

	Boards []Board `json:"boards"`
}
