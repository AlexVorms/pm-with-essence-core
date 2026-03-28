package repository

import (
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectStorage struct {
	db *gorm.DB
}

func NewProjectStorage(db *gorm.DB) *ProjectStorage {
	return &ProjectStorage{
		db: db,
	}
}
func (p *ProjectStorage) CreateProject(project pm_entity.Project) error {
	err := p.db.Create(&project).Error
	if err != nil {
		return err
	}
	return nil
}
func (p *ProjectStorage) GetProject(projectId uuid.UUID) (error, *pm_entity.Project) {
	var project pm_entity.Project
	err := p.db.Table("projects").First(&project, projectId).Error
	if err != nil {
		return err, nil
	}
	return nil, &project
}
func (p *ProjectStorage) DeleteProject(projectID uuid.UUID) error {
	err := p.db.Where("project_id = ?", projectID).Delete(&pm_entity.Project{}).Error
	return err
}
func (p *ProjectStorage) UpdateProject(project pm_entity.Project) error {
	err := p.db.Model(&project).Updates(project).Error
	return err
}
