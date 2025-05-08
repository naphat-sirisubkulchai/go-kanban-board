package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/config"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
)

type BoardRepository interface {
	CreateBoard(board *models.Board) error
	UpdateBoardTitle(boardID, title string) error
	DeleteBoard(boardID string) error
	AddMember(boardID, userID string) error
	GetBoardByName(name string) (*models.Board, error)
}

type boardRepo struct{}

func NewBoardRepository() BoardRepository {
	return &boardRepo{}
}

func (r *boardRepo) CreateBoard(board *models.Board) error {
	return config.DB.Create(board).Error
}

func (r *boardRepo) UpdateBoardTitle(boardID, title string) error {
	return config.DB.Model(&models.Board{}).Where("id = ?", boardID).Update("title", title).Error
}

func (r *boardRepo) DeleteBoard(boardID string) error {
	return config.DB.Delete(&models.Board{}, "id = ?", boardID).Error
}
func (r *boardRepo) AddMember(boardID, userID string) error {
	var board models.Board
	if err := config.DB.Preload("Members").First(&board, "id = ?", boardID).Error; err != nil {
		return err
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		return err
	}

	return config.DB.Model(&board).Association("Members").Append(&user)
}
func (r *boardRepo) GetBoardByName(name string) (*models.Board, error) {
	var board models.Board
	err := config.DB.Preload("Members").Where("name = ?", name).First(&board).Error
	if err != nil {
		return nil, err
	}
	return &board, nil
}