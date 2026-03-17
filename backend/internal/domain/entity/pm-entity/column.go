package pm_entity

import "github.com/google/uuid"

type Column struct {
	ID      uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name    string    `json:"name"`
	IsFinal bool      `json:"is_final"`
	Order   int       `json:"order"` //порядок начинается с 0

	BoardID uuid.UUID `json:"boardID" gorm:"type:uuid"`
	Board   *Board    `json:"board"`
	Issues  []Issue   `json:"issues"`
}
