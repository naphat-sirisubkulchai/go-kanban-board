package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/config"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
)

type ColumnRepository interface {
	Create(column *models.Column) error
	UpdateName(columnID, newName string) error
	Delete(columnID string) error
	GetByBoardID(boardID string) ([]models.Column, error)
}

type columnRepo struct{}

func NewColumnRepository() ColumnRepository {
	return &columnRepo{}
}

func (r *columnRepo) Create(column *models.Column) error {
	return config.DB.Create(column).Error
}

func (r *columnRepo) UpdateName(columnID, newName string) error {
	return config.DB.Model(&models.Column{}).
		Where("id = ?", columnID).
		Update("name", newName).Error
}

func (r *columnRepo) Delete(columnID string) error {
	return config.DB.Delete(&models.Column{}, "id = ?", columnID).Error
}

func (r *columnRepo) GetByBoardID(boardID string) ([]models.Column, error) {
	var cols []models.Column
	err := config.DB.Where("board_id = ?", boardID).Order("position ASC").Find(&cols).Error
	return cols, err
}
