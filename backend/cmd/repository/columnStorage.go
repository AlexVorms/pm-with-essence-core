package repository

import (
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ColumnStorage struct {
	db *gorm.DB
}

func NewColumnStorage(db *gorm.DB) *ColumnStorage {
	return &ColumnStorage{
		db: db,
	}
}
func (c *ColumnStorage) Create(column *pm_entity.Column) error {
	err := c.db.Create(&column).Error
	return err
}
func (c *ColumnStorage) GetColumn(columnID uuid.UUID) (*pm_entity.Column, error) {
	var column pm_entity.Column
	err := c.db.Model(pm_entity.Column{}).Find(&column, columnID).Error
	return &column, err
}
func (c *ColumnStorage) UpdateColumn(column *pm_entity.Column) error {
	err := c.db.Model(pm_entity.Column{}).Updates(column).Error
	return err
}
func (c *ColumnStorage) DeleteColumn(columnID uuid.UUID) error {
	err := c.db.Delete(pm_entity.Column{}, columnID).Error
	return err
}
