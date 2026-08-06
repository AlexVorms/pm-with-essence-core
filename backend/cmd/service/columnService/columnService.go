package columnService

import (
	"pm-with-essence/cmd/api/model/columnDTO"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
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
func (p *ColumnService) CreateColumn(boardID uuid.UUID, model columnDTO.ColumnDTO) *app_errors.HttpError {
	_, err := p.boardStorage.GetBoard(boardID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This board doesn't exist")
	}

	newColumn := pm_entity.CreateColumnEntity(model.Name, model.IsFinal, model.Order, boardID)
	err1 := p.columnStorage.Create(newColumn)
	if err1 != nil {
		return app_errors.NewPostgresWriteError(err1, "Error while creating column")
	}
	return nil
}
func (c *ColumnService) DeleteColumn(columnID uuid.UUID) *app_errors.HttpError {
	_, err := c.columnStorage.GetColumn(columnID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This column doesn't exist")
	}
	err = c.columnStorage.DeleteColumn(columnID)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while deleting column")
	}
	return nil
}
func (c *ColumnService) UpdateColumn(columnID uuid.UUID, model columnDTO.ColumnDTO) *app_errors.HttpError {
	column, err := c.columnStorage.GetColumn(columnID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This column doesn't exist")
	}
	column.IsFinal = model.IsFinal
	column.Name = model.Name
	err = c.columnStorage.UpdateColumn(column)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while updating column")
	}
	return nil
}
func (c *ColumnService) ChangeColumnOrder(model columnDTO.UpdateColumnOrderDTO) *app_errors.HttpError {
	firstColumn, err := c.columnStorage.GetColumn(model.FirstColumnID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "First column doesn't exist")
	}
	secondColumn, err1 := c.columnStorage.GetColumn(model.SecondColumnID)
	if err1 != nil {
		return app_errors.NewNotFoundError(err1, "Second column doesn't exist")
	}
	firstColumn.Order = model.FirstOrder
	secondColumn.Order = model.SecondOrder

	err = c.columnStorage.UpdateColumn(firstColumn)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while updating column")
	}
	err = c.columnStorage.UpdateColumn(secondColumn)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while updating column")
	}
	return nil
}
