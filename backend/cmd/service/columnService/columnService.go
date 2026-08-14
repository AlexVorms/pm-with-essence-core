package columnService

import (
	"errors"
	"fmt"
	"pm-with-essence/cmd/api/model/columnDTO"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ColumnService struct {
	boardStorage  *repository.BoardStorage
	columnStorage *repository.ColumnStorage
}

func NewColumnService(columnStorage *repository.ColumnStorage, storage *repository.BoardStorage) *ColumnService {
	return &ColumnService{
		columnStorage: columnStorage,
		boardStorage:  storage,
	}
}
func (c *ColumnService) CreateColumn(boardID uuid.UUID, model columnDTO.ColumnDTO) *app_errors.HttpError {
	board, err := c.boardStorage.GetBoard(boardID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read board from database",
			)
		}
		return app_errors.NewNotFoundError(err, "This board doesn't exist")
	}

	newColumn := pm_entity.CreateColumnEntity(model.Name, false, c.getNextColumnOrder(board), boardID)
	fmt.Println(newColumn)
	err1 := c.columnStorage.Create(newColumn)
	if err1 != nil {
		return app_errors.NewPostgresWriteError(err1, "Error while creating column")
	}

	return nil
}
func (p *ColumnService) getNextColumnOrder(board *pm_entity.Board) int {
	if len(board.Columns) == 0 {
		return 0
	}

	maxOrder := board.Columns[0].Order

	for _, column := range board.Columns[1:] {
		if column.Order > maxOrder {
			maxOrder = column.Order
		}
	}
	return maxOrder + 1
}
func (p *ColumnService) DeleteColumn(columnID uuid.UUID) *app_errors.HttpError {

	_, err := p.columnStorage.GetColumn(columnID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read column from database",
			)
		}
		return app_errors.NewNotFoundError(err, "This column doesn't exist")
	}

	err = p.columnStorage.DeleteColumn(columnID)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while deleting column")
	}
	return nil
}
func (p *ColumnService) UpdateColumn(columnID uuid.UUID, model columnDTO.ColumnUpdateDTO) *app_errors.HttpError {
	column, err := p.columnStorage.GetColumn(columnID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read column from database",
			)
		}
		return app_errors.NewNotFoundError(err, "This column doesn't exist")
	}

	column.IsFinal = model.IsFinal
	column.Name = model.Name

	err = p.columnStorage.UpdateColumn(column)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while updating column")
	}

	return nil
}
func (p *ColumnService) ChangeColumnOrder(model columnDTO.UpdateColumnOrderDTO) *app_errors.HttpError {
	firstColumn, err := p.columnStorage.GetColumn(model.FirstColumnID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read column from database",
			)
		}
		return app_errors.NewNotFoundError(err, "First column doesn't exist")
	}

	secondColumn, err1 := p.columnStorage.GetColumn(model.SecondColumnID)
	if err1 != nil {
		if !errors.Is(err1, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err1,
				"Failed to read column from database",
			)
		}
		return app_errors.NewNotFoundError(err1, "Second column doesn't exist")
	}

	firstColumn.Order = model.FirstOrder
	secondColumn.Order = model.SecondOrder

	err = p.columnStorage.UpdateColumn(firstColumn)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while updating column")
	}

	err = p.columnStorage.UpdateColumn(secondColumn)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while updating column")
	}
	return nil
}
