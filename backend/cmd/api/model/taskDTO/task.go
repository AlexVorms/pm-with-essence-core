package taskDTO

import (
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"
	"time"

	"github.com/google/uuid"
)

type TaskDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type GetTaskDTO struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt" `
	UpdatedAt   time.Time `json:"updateDate"`
	CompletedAt time.Time `json:"completedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`
}

func CreateNewGetTaskDTO(issue *pm_entity.Issue) *GetTaskDTO {
	return &GetTaskDTO{
		ID:          issue.ID,
		CreatedAt:   issue.CreatedAt,
		UpdatedAt:   issue.UpdatedAt,
		CompletedAt: issue.CompletedAt,
		Name:        issue.Name,
		Description: issue.Description,
		IsComplete:  issue.IsComplete,
	}
}
