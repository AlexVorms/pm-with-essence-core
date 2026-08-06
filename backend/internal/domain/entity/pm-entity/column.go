package pm_entity

import "github.com/google/uuid"

type ColumnData struct {
	ID      uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name    string    `json:"name"`
	IsFinal bool      `json:"is_final"`
	Order   int       `json:"order"` //порядок начинается с 0

	BoardID uuid.UUID `json:"boardID" gorm:"type:uuid"`
	Board   *Board    `json:"board"`
	Issues  []Issue   `json:"issues"`
}

func CreateColumnEntity(Name string, IsFinal bool, Order int, boardID uuid.UUID) *ColumnData {
	return &ColumnData{
		ID:      uuid.New(),
		Name:    Name,
		IsFinal: IsFinal,
		Order:   Order,
		BoardID: boardID,
	}
}
