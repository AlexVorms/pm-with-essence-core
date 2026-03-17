package pm_entity

import "github.com/google/uuid"

type Issue struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsComplete  bool      `json:"isComplete"`

	Owner uuid.UUID `json:"owner" gorm:"type:uuid"`

	ColumnID uuid.UUID `json:"columnID" gorm:"type:uuid"`
	Column   *Column   `json:"column"`
}
