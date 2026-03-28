package columnService

import (
	"pm-with-essence/cmd/api/model/columnDTO"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/internal/api/errors"
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
func (p *ColumnService) CreateColumn(boardID uuid.UUID, model columnDTO.ColumnDTO) *errors.HttpError {
	_, err := p.boardStorage.GetBoard(boardID)
	if err != nil {
		return errors.NewNotFoundError(err, "This board doesn't exist")
	}

	newColumn := pm_entity.CreateColumnEntity(model.Name, model.IsFinal, model.Order, boardID)
	err1 := p.columnStorage.Create(newColumn)
	if err1 != nil {
		return errors.NewPostgresWriteError(err1, "Error while creating column")
	}
	return nil
}
