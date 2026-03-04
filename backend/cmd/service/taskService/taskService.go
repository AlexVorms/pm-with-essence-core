package taskService

import "pm-with-essence/cmd/repository"

type TaskService struct {
	taskStorage *repository.TaskStorage
}

func NewTaskService(taskStorage *repository.TaskStorage) *TaskService {
	return &TaskService{
		taskStorage: taskStorage,
	}
}
