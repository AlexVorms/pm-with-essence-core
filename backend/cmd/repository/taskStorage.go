package repository

import (
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskStorage struct {
	db *gorm.DB
}

func NewTaskStorage(db *gorm.DB) *TaskStorage {
	return &TaskStorage{
		db: db,
	}
}
func (s *TaskStorage) CreateTask(task pm_entity.Issue) error {
	return s.db.Model(pm_entity.Issue{}).Create(&task).Error
}
func (s *TaskStorage) GetTask(taskID uuid.UUID) (*pm_entity.Issue, error) {
	var task pm_entity.Issue
	err := s.db.Model(pm_entity.Issue{}).First(&task, taskID).Error
	return &task, err
}
func (s *TaskStorage) UpdateTask(task *pm_entity.Issue) error {
	return s.db.Model(pm_entity.Issue{}).Updates(&task).Error
}
func (s *TaskStorage) DeleteTask(taskID uuid.UUID) error {
	err := s.db.Delete(&pm_entity.Issue{}, taskID).Error
	return err
}
