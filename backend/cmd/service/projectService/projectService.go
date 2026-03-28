package projectService

import (
	"pm-with-essence/cmd/api/model/projectDTO"
	"pm-with-essence/cmd/repository"
	"pm-with-essence/internal/api/errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
)

type ProjectService struct {
	projectStorage *repository.ProjectStorage
}

func NewProjectService(projectStorage *repository.ProjectStorage) *ProjectService {
	return &ProjectService{
		projectStorage: projectStorage,
	}
}
func (p *ProjectService) CreateProject(model projectDTO.ProjectDTO) *errors.HttpError {
	project := pm_entity.Project{
		ID:          uuid.New(),
		Name:        model.Name,
		Description: model.Description,
		IsPublic:    model.IsPublic,
	}
	err := p.projectStorage.CreateProject(project)
	if err != nil {
		return errors.NewPostgresWriteError(err, "faled database to create project")
	}
	return nil
}
func (p *ProjectService) DeleteProject(projectId uuid.UUID) *errors.HttpError {
	err, _ := p.projectStorage.GetProject(projectId)
	if err != nil {
		return errors.NewNotFoundError(err, "this project doesn't exist")
	}
	err = p.projectStorage.DeleteProject(projectId)
	if err != nil {
		return errors.NewPostgresWriteError(err, "faled to delete the project")
	}
	return nil
}
func (p *ProjectService) GetProject(projectId uuid.UUID) (*pm_entity.Project, *errors.HttpError) {
	err, project := p.projectStorage.GetProject(projectId)
	if err != nil {
		return nil, errors.NewPostgresReadError(err, "this project doesn't exist")
	}
	return project, nil
}
func (p *ProjectService) GetAllProjects(model projectDTO.ProjectDTO) *errors.HttpError {
	return nil
}
func (p *ProjectService) UpdateProject(projectID uuid.UUID, model projectDTO.ProjectDTO) *errors.HttpError {
	err, _ := p.projectStorage.GetProject(projectID)
	if err != nil {
		return errors.NewNotFoundError(err, "this project doesn't exist")
	}
	project := pm_entity.Project{
		ID:          projectID,
		Name:        model.Name,
		Description: model.Description,
		IsPublic:    model.IsPublic,
	}
	err = p.projectStorage.UpdateProject(project)
	if err != nil {
		return errors.NewPostgresWriteError(err, "faled to update the project")
	}
	return nil
}
func (p *ProjectService) IsProjectExist(projectId uuid.UUID) (bool, error) {
	err, _ := p.projectStorage.GetProject(projectId)
	if err != nil {
		return false, err
	}
	return true, nil
}
