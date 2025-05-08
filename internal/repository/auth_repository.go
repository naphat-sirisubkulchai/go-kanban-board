package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/config"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type authRepo struct{}

func NewAuthRepository() AuthRepository {
	return &authRepo{}
}

func (r *authRepo) CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *authRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

