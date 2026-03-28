package repository

import (
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

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
