package repository

import "gorm.io/gorm"

type ProjectStorage struct {
	db *gorm.DB
}

func NewProjectStorage(db *gorm.DB) *ProjectStorage {
	return &ProjectStorage{
		db: db,
	}
}
