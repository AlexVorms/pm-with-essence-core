package columnDTO

import "github.com/google/uuid"

type ColumnDTO struct {
	Name    string `json:"name"`
	IsFinal bool   `json:"is_final"`
	Order   int    `json:"order"` //порядок начинается с 0
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
