package boardService

import (
	"pm-with-essence/cmd/api/model/boardDTO"
	"pm-with-essence/cmd/api/model/columnDTO"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/cmd/service/columnService"
	"pm-with-essence/cmd/service/projectService"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
)

type BoardService struct {
	boardStorage   *repository.BoardStorage
	columnService  *columnService.ColumnService
	projectService *projectService.ProjectService
}

func NewBoardService(boardStorage *repository.BoardStorage,
	service *columnService.ColumnService,
	projectService *projectService.ProjectService) *BoardService {
	return &BoardService{
		boardStorage:   boardStorage,
		columnService:  service,
		projectService: projectService,
	}
}
func (b *BoardService) CreateBoard(projectID uuid.UUID, model boardDTO.BoardDTO) *app_errors.HttpError {
	isProjectExist, err := b.projectService.IsProjectExist(projectID)
	if !isProjectExist {
		return app_errors.NewNotFoundError(err, "This project doesn't exist")
	}

	board := pm_entity.CreateBoardEntity(model.Name, model.Description, model.IsPublic, projectID)
	err = b.boardStorage.CreateBoard(board)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error creating a board in the database")
	}

	column := columnDTO.ColumnDTO{
		Name:    "Новая колонка",
		IsFinal: false,
		Order:   0,
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
		return app_errors.NewNotFoundError(err, "This board doesn't exist")
	}
	err1 := b.boardStorage.DeleteBoard(boardID)
	if err1 != nil {
		return app_errors.NewPostgresWriteError(err1, "Error deleting a board")
	}
	return nil
}
func (b *BoardService) GetBoard(boardID uuid.UUID) (*pm_entity.Board, *app_errors.HttpError) {
	//TODO:рефакторинг, сделать проверку на существование и получение доски одним запросом
	isBoardExist, err := b.isBoardExist(boardID)
	if !isBoardExist {
		return nil, app_errors.NewNotFoundError(err, "This board doesn't exist")
	}

	board, err1 := b.boardStorage.GetBoard(boardID)
	if err1 != nil {
		return nil, app_errors.NewPostgresReadError(err1, "Error getting a board")
	}

	return board, nil
}
func (b *BoardService) GetAllBoards(projectID uuid.UUID) ([]*pm_entity.Board, *app_errors.HttpError) {

	isProjectExist, err := b.projectService.IsProjectExist(projectID)
	if !isProjectExist {
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
