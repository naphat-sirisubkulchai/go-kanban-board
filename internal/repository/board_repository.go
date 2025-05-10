package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"gorm.io/gorm"
)

type BoardRepository interface {
	CreateBoard(board *models.Board) error
	UpdateBoardTitle(boardID, title string) error
	DeleteBoard(boardID string) error
	AddMember(boardID, userID string) error
	GetBoardByName(name string) (*models.Board, error)
}

type boardRepo struct {
	db *gorm.DB
}

func NewBoardRepository(db *gorm.DB) BoardRepository {
	return &boardRepo{db: db}
}

func (r *boardRepo) CreateBoard(board *models.Board) error {
	return r.db.Create(board).Error
}

func (r *boardRepo) UpdateBoardTitle(boardID, title string) error {
	return r.db.Model(&models.Board{}).Where("id = ?", boardID).Update("title", title).Error
}

func (r *boardRepo) DeleteBoard(boardID string) error {
	return r.db.Delete(&models.Board{}, "id = ?", boardID).Error
}

func (r *boardRepo) AddMember(boardID, userID string) error {
	var board models.Board
	if err := r.db.Preload("Members").First(&board, "id = ?", boardID).Error; err != nil {
		return err
	}

	var user models.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return err
	}

	return r.db.Model(&board).Association("Members").Append(&user)
}

func (r *boardRepo) GetBoardByName(name string) (*models.Board, error) {
	var board models.Board
	err := r.db.Preload("Members").Where("name = ?", name).First(&board).Error
	if err != nil {
		return nil, err
	}
	return &board, nil
}
