package repository

import "gorm.io/gorm"

type TaskStorage struct {
	db *gorm.DB
}

func NewTaskStorage(db *gorm.DB) *TaskStorage {
	return &TaskStorage{
		db: db,
	}
}
