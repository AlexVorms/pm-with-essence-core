package projectService

import (
	"errors"
	"pm-with-essence/cmd/api/model/projectDTO"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectService struct {
	projectStorage ProjectStorage
}
type ProjectStorage interface {
	CreateProject(project pm_entity.Project) error
	GetProject(projectId uuid.UUID) (error, *pm_entity.Project)
	DeleteProject(projectID uuid.UUID) error
	UpdateProject(project pm_entity.Project) error
}

func NewProjectService(projectStorage ProjectStorage) *ProjectService {
	return &ProjectService{
		projectStorage: projectStorage,
	}
}
func (p *ProjectService) CreateProject(model projectDTO.ProjectDTO) *app_errors.HttpError {
	project := pm_entity.Project{
		ID:          uuid.New(),
		Name:        model.Name,
		Description: model.Description,
		IsPublic:    model.IsPublic,
	}
	err := p.projectStorage.CreateProject(project)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "faled database to create project")
	}
	return nil
}
func (p *ProjectService) DeleteProject(projectId uuid.UUID) *app_errors.HttpError {
	err, _ := p.projectStorage.GetProject(projectId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read project from database",
			)
		}
		return app_errors.NewNotFoundError(err, "This project doesn't exist")
	}
	err = p.projectStorage.DeleteProject(projectId)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "faled to delete the project")
	}
	return nil
}
func (p *ProjectService) GetProject(projectId uuid.UUID) (*pm_entity.Project, *app_errors.HttpError) {
	err, project := p.projectStorage.GetProject(projectId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_errors.NewPostgresReadError(
				err,
				"Failed to read project from database",
			)
		}
		return nil, app_errors.NewPostgresReadError(err, "this project doesn't exist")
	}
	return project, nil
}
func (p *ProjectService) GetAllProjects(model projectDTO.ProjectDTO) *app_errors.HttpError {
	return nil
}
func (p *ProjectService) UpdateProject(projectID uuid.UUID, model projectDTO.ProjectDTO) *app_errors.HttpError {
	err, _ := p.projectStorage.GetProject(projectID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return app_errors.NewPostgresReadError(
				err,
				"Failed to read project from database",
			)
		}
		return app_errors.NewNotFoundError(err, "this project doesn't exist")
	}
	project := pm_entity.Project{
		ID:          projectID,
		Name:        model.Name,
		Description: model.Description,
		IsPublic:    model.IsPublic,
	}
	err = p.projectStorage.UpdateProject(project)
	if err != nil {
		return app_errors.NewPostgresWriteError(err, "faled to update the project")
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
