package boardDTO

import (
	"pm-with-essence/cmd/api/model/columnDTO"

	"github.com/google/uuid"
)

type BoardDTO struct {
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	IsPublic    bool   `json:"isPublic"`
}
type GetBoardDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"isPublic"`

	Columns []columnDTO.GetColumnDTO `json:"columns"`
}
