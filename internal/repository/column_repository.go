package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"gorm.io/gorm"
)

type ColumnRepository interface {
	Create(column *models.Column) error
	UpdateName(columnID, newName string) error
	Delete(columnID string) error
	GetByBoardID(boardID string) ([]models.Column, error)
}

type columnRepo struct {
	db *gorm.DB
}

func NewColumnRepository(db *gorm.DB) ColumnRepository {
	return &columnRepo{db: db}
}

func (r *columnRepo) Create(column *models.Column) error {
	return r.db.Create(column).Error
}

func (r *columnRepo) UpdateName(columnID, newName string) error {
	return r.db.Model(&models.Column{}).
		Where("id = ?", columnID).
		Update("name", newName).Error
}

func (r *columnRepo) Delete(columnID string) error {
	return r.db.Delete(&models.Column{}, "id = ?", columnID).Error
}

func (r *columnRepo) GetByBoardID(boardID string) ([]models.Column, error) {
	var cols []models.Column
	err := r.db.Where("board_id = ?", boardID).Order("position ASC").Find(&cols).Error
	return cols, err
}
