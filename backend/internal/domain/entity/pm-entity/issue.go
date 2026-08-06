package pm_entity

import (
	"time"

	"github.com/google/uuid"
)

type Issue struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt   time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt   time.Time `json:"updateDate"`
	CompletedAt time.Time `json:"completedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
	IsComplete  bool   `json:"isComplete"`

	Owner uuid.UUID `json:"owner" gorm:"type:uuid"`

	ColumnID uuid.UUID `json:"columnID" gorm:"type:uuid"`
	ColumnData   *ColumnData   `json:"column"`
}

func CreateNewIssueEntity(Name string, Description string, ColumnID uuid.UUID) Issue {
	return Issue{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Name:        Name,
		Description: Description,
		IsComplete:  false,
		ColumnID:    ColumnID,
	}
}
func CompleteIssueEntity(issue *Issue) *Issue {
	issue.IsComplete = true
	issue.UpdatedAt = time.Now()
	issue.CompletedAt = time.Now()
	return issue
}
func UpdateIssueEntity(issue *Issue, Name string, Description string) *Issue {
	issue.UpdatedAt = time.Now()
	issue.Name = Name
	issue.Description = Description
	return issue
}
