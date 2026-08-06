package taskService

import (
	"pm-with-essence/cmd/api/model/taskDTO"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
)

type TaskService struct {
	taskStorage   *repository.TaskStorage
	columnStorage *repository.ColumnStorage
}

func NewTaskService(taskStorage *repository.TaskStorage,
	columnStorage *repository.ColumnStorage) *TaskService {
	return &TaskService{
		taskStorage:   taskStorage,
		columnStorage: columnStorage,
	}
}
func (ts *TaskService) CreateTask(columnID uuid.UUID, model taskDTO.TaskDTO) *app_errors.HttpError {
	_, err := ts.columnStorage.GetColumn(columnID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This column doesn't exist")
	}
	task := pm_entity.CreateNewIssueEntity(model.Name, model.Description, columnID)
	err = ts.taskStorage.CreateTask(task)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while creating task")
	}
	return nil
}
func (ts *TaskService) CompleteTask(taskID uuid.UUID) *app_errors.HttpError {
	task, err := ts.taskStorage.GetTask(taskID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This task doesn't exist")
	}
	//TODO:Если есть проблемы, они здесь
	task = pm_entity.CompleteIssueEntity(task)
	err = ts.taskStorage.UpdateTask(task)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while completing task")
	}
	return nil
}
func (ts *TaskService) UpdateTask(taskID uuid.UUID, model taskDTO.TaskDTO) *app_errors.HttpError {
	task, err := ts.taskStorage.GetTask(taskID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This task doesn't exist")
	}
	issue := pm_entity.UpdateIssueEntity(task, model.Name, model.Description)
	err = ts.taskStorage.UpdateTask(issue)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while updating task")
	}
	return nil
}
func (ts *TaskService) DeleteTask(taskID uuid.UUID) *app_errors.HttpError {
	_, err := ts.taskStorage.GetTask(taskID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This task doesn't exist")
	}
	err = ts.taskStorage.DeleteTask(taskID)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while deleting task")
	}
	return nil
}
func (ts *TaskService) GetTask(taskID uuid.UUID) (*taskDTO.GetTaskDTO, *app_errors.HttpError) {
	task, err := ts.taskStorage.GetTask(taskID)
	if err != nil {
		return nil, app_errors.NewNotFoundError(err, "This task doesn't exist")
	}
	newTaskDTO := taskDTO.CreateNewGetTaskDTO(task)
	return newTaskDTO, nil
}
func (ts *TaskService) ChangeTaskColumn(taskID uuid.UUID, columnID uuid.UUID) *app_errors.HttpError {
	task, err := ts.taskStorage.GetTask(taskID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This task doesn't exist")
	}
	_, err = ts.columnStorage.GetColumn(columnID)
	if err != nil {
		return app_errors.NewNotFoundError(err, "This column doesn't exist")
	}
	task.ColumnID = columnID
	err = ts.taskStorage.UpdateTask(task)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "Error while updating task")
	}
	return nil
}
