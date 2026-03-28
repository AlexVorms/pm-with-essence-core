package pm_entity

import "github.com/google/uuid"

type Board struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"isPublic"`

	ProjectID uuid.UUID `json:"projectID" gorm:"type:uuid"`
	Project   *Project  `json:"project"`

	Columns []Column `json:"columns" gorm:"constraint:OnDelete:CASCADE;"`
	//Status      string    `json:"status" gorm:"not null"`
}

func CreateBoardEntity(Name string, Description string, IsPublic bool, ProjectID uuid.UUID) *Board {
	return &Board{
		ID:          uuid.New(),
		Name:        Name,
		Description: Description,
		IsPublic:    IsPublic,
		ProjectID:   ProjectID,
	}
}
