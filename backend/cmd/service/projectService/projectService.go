package projectService

import "pm-with-essence/cmd/repository"

type ProjectService struct {
	projectStorage *repository.ProjectStorage
}

func NewProjectService(projectStorage *repository.ProjectStorage) *ProjectService {
	return &ProjectService{
		projectStorage: projectStorage,
	}
}
