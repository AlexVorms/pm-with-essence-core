package repository

import (
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BoardStorage struct {
	db *gorm.DB
}

func NewBoardStorage(db *gorm.DB) *BoardStorage {
	return &BoardStorage{
		db: db,
	}
}
func (b *BoardStorage) GetBoard(boardID uuid.UUID) (*pm_entity.Board, error) {
	var board *pm_entity.Board
	err := b.db.Model(pm_entity.Board{}).Preload("Columns").Preload("Columns.Issues").Find(&board, boardID).Error
	if err != nil {
		return nil, err
	}
	return board, nil
}
func (b *BoardStorage) GetProjectBoards(projectID uuid.UUID) ([]*pm_entity.Board, error) {
	var boards []*pm_entity.Board
	err := b.db.Where(pm_entity.Board{ProjectID: projectID}).Find(&boards).Error
	if err != nil {
		return nil, err
	}
	return boards, nil
}
func (b *BoardStorage) DeleteBoard(boardID uuid.UUID) error {
	err := b.db.Delete(&pm_entity.Board{}, boardID).Error
	return err
}
func (b *BoardStorage) CreateBoard(board *pm_entity.Board) error {
	err := b.db.Create(board).Error
	if err != nil {
		return err
	}
	return nil
}
func (b *BoardStorage) UpdateBoard(board *pm_entity.Board) error {
	err := b.db.Model(&board).Updates(board).Error
	return err
}
