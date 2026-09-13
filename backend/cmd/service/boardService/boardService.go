package boardService

import (
	"errors"
	"fmt"
	"pm-with-essence/cmd/api/model/boardDTO"
	"pm-with-essence/cmd/api/model/columnDTO"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"
	"pm-with-essence/internal/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardService struct {
	boardStorage   BoardStorage
	columnService  ColumnService
	projectService ProjectService
}
type BoardStorage interface {
	GetBoard(boardID uuid.UUID) (*pm_entity.Board, error)
	GetProjectBoards(projectID uuid.UUID) ([]*pm_entity.Board, error)
	DeleteBoard(boardID uuid.UUID) error
	CreateBoard(board *pm_entity.Board) error
	UpdateBoard(board *pm_entity.Board) error
}
type ColumnService interface {
	CreateColumn(boardID uuid.UUID, model columnDTO.ColumnDTO) *app_errors.HttpError
}
type ProjectService interface {
	GetProject(projectId uuid.UUID) (*pm_entity.Project, *app_errors.HttpError)
	IsProjectExist(projectId uuid.UUID) (bool, error)
}

func NewBoardService(boardStorage BoardStorage,
	service ColumnService,
	projectService ProjectService) *BoardService {
	return &BoardService{
		boardStorage:   boardStorage,
		columnService:  service,
		projectService: projectService,
	}
}
func (b *BoardService) CreateBoard(projectID uuid.UUID, model boardDTO.BoardDTO) *app_errors.HttpError {
	isProjectExist, err := b.projectService.IsProjectExist(projectID)
	if !isProjectExist {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read project from database",
			)
		}
		return app_errors.NewNotFoundError(err, "This project doesn't exist")
	}

	board := pm_entity.CreateBoardEntity(model.Name, model.Description, model.IsPublic, projectID)

	err = b.boardStorage.CreateBoard(board)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error creating a board in the database")
	}

	column := columnDTO.ColumnDTO{
		Name: "Новая колонка",
	}
	err1 := b.columnService.CreateColumn(board.ID, column)
	if err1 != nil {
		return err1
	}
	return nil
}
func (b *BoardService) DeleteBoard(boardID uuid.UUID) *app_errors.HttpError {
	isBoardExist, err := b.isBoardExist(boardID)
	if !isBoardExist {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read board from database",
			)
		}
		return app_errors.NewNotFoundError(err, "This board doesn't exist")
	}
	err1 := b.boardStorage.DeleteBoard(boardID)
	if err1 != nil {
		return app_errors.NewPostgresWriteError(err1, "Error deleting a board")
	}
	return nil
}
func (b *BoardService) GetBoard(boardID uuid.UUID) (*boardDTO.GetBoardDTO, *app_errors.HttpError) {

	board, err := b.boardStorage.GetBoard(boardID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_errors.NewPostgresReadError(
				err,
				"Failed to read board from database",
			)
		}
		return nil, app_errors.NewNotFoundError(err, "This board doesn't exist")
	}
	fmt.Println(board)

	return mapper.ToBoardResponse(board), nil
}
func (b *BoardService) GetAllBoards(projectID uuid.UUID) ([]*pm_entity.Board, *app_errors.HttpError) {

	isProjectExist, err := b.projectService.IsProjectExist(projectID)
	if !isProjectExist {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_errors.NewPostgresReadError(
				err,
				"Failed to read project from database",
			)
		}
		return nil, app_errors.NewNotFoundError(err, "This project doesn't exist")
	}

	projects, err1 := b.boardStorage.GetProjectBoards(projectID)
	if err1 != nil {
		return nil, app_errors.NewPostgresReadError(err1, "Error getting a boards")
	}

	return projects, nil
}
func (b *BoardService) UpdateBoard(boardID uuid.UUID, model boardDTO.BoardDTO) *app_errors.HttpError {
	board, err := b.boardStorage.GetBoard(boardID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read board from database",
			)
		}
		return app_errors.NewNotFoundError(err, "This board doesn't exist")
	}
	board.Name = model.Name
	board.Description = model.Description
	board.IsPublic = model.IsPublic
	err = b.boardStorage.UpdateBoard(board)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error updating a board")
	}
	return nil
}
func (b *BoardService) isBoardExist(boardID uuid.UUID) (bool, error) {
	_, err := b.boardStorage.GetBoard(boardID)
	if err != nil {
		return false, err
	}
	return true, nil
}
