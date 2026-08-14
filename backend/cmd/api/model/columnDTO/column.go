package columnDTO

import (
	"pm-with-essence/cmd/api/model/taskDTO"

	"github.com/google/uuid"
)

type ColumnDTO struct {
	Name    string `json:"name"`
	IsFinal bool   `json:"is_final"`
	Order   int    `json:"order"` //порядок начинается с 0
}

type GetColumnDTO struct {
	ID      uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name    string    `json:"name"`
	IsFinal bool      `json:"is_final"`
	Order   int       `json:"order"` //порядок начинается с 0

	Tasks []taskDTO.GetColumnTasksDTO `json:"tasks"`
}
type ColumnUpdateDTO struct {
	Name    string `json:"name"`
	IsFinal bool   `json:"is_final"`
}
type UpdateColumnOrderDTO struct {
	FirstColumnID  uuid.UUID `json:"first_column_id"`
	FirstOrder     int       `json:"first_order"`
	SecondColumnID uuid.UUID `json:"second_column_id"`
	SecondOrder    int       `json:"second_order"`
}
