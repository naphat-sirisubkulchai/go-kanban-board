package usecase

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUsecase) GetUserByEmail(email string) (*models.User, error) {
	return u.repo.GetUserByEmail(email)
}
