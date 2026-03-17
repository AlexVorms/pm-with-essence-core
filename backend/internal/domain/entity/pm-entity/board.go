package pm_entity

import "github.com/google/uuid"

type Board struct {
	ID             uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name           string    `json:"name" gorm:"not null"`
	Description    string    `json:"description"`
	IsPublic       bool      `json:"isPublic"`
	NumberOfColumn int       `json:"numberOfColumn"`

	ProjectID uuid.UUID `json:"projectID" gorm:"type:uuid"`
	Project   *Project  `json:"project"`

	Columns []Column `json:"columns"`
	//Status      string    `json:"status" gorm:"not null"`
}
