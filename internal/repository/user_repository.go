package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/config"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
