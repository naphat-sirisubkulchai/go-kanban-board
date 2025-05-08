package usecase

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/repository"
)

type ColumnUsecase interface {
	Create(column *models.Column) error
	UpdateName(columnID, newName string) error
	Delete(columnID string) error
	GetByBoardID(boardID string) ([]models.Column, error)
}

type columnUsecase struct {
	repo repository.ColumnRepository
}

func NewColumnUsecase(r repository.ColumnRepository) ColumnUsecase {
	return &columnUsecase{repo: r}
}

func (u *columnUsecase) Create(column *models.Column) error {
	return u.repo.Create(column)
}

func (u *columnUsecase) UpdateName(columnID, newName string) error {
	return u.repo.UpdateName(columnID, newName)
}

func (u *columnUsecase) Delete(columnID string) error {
	return u.repo.Delete(columnID)
}

func (u *columnUsecase) GetByBoardID(boardID string) ([]models.Column, error) {
	return u.repo.GetByBoardID(boardID)
}
