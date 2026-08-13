package mapper

import (
	"pm-with-essence/cmd/api/model/boardDTO"
	"pm-with-essence/cmd/api/model/columnDTO"
	"pm-with-essence/cmd/api/model/taskDTO"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"
)

func ToBoardResponse(board *pm_entity.Board) *boardDTO.GetBoardDTO {
	return &boardDTO.GetBoardDTO{
		ID:          board.ID,
		Name:        board.Name,
		Description: board.Description,
		IsPublic:    board.IsPublic,
		ProjectID:   board.ProjectID,
		Columns:     ToColumnResponses(board.Columns),
	}
}
func ToColumnResponses(columns []pm_entity.Column) []columnDTO.GetColumnDTO {
	result := make([]columnDTO.GetColumnDTO, 0, len(columns))

	for _, column := range columns {
		result = append(result, columnDTO.GetColumnDTO{
			ID:      column.ID,
			Name:    column.Name,
			IsFinal: column.IsFinal,
			Order:   column.Order,
			Tasks:   ToIssueResponses(column.Issues),
		})
	}

	return result
}
func ToIssueResponses(issues []pm_entity.Issue) []taskDTO.GetColumnTasksDTO {
	result := make([]taskDTO.GetColumnTasksDTO, 0, len(issues))

	for _, issue := range issues {
		result = append(result, taskDTO.GetColumnTasksDTO{
			ID:         issue.ID,
			Name:       issue.Name,
			IsComplete: issue.IsComplete,
		})
	}

	return result
}
