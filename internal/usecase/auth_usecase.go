package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/repository"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/utils"
)

type AuthUsecase interface {
	Register(user *models.User) error
	Login(email, password string) (string, error)
}

type authUsecase struct {
	repo repository.AuthRepository
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &authUsecase{repo: repo}
}

func (u *authUsecase) Register(user *models.User) error {
	if !utils.IsEmailValid(user.Email) {
		return errors.New("invalid email format")
	}

	existingUser, _ := u.repo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	return u.repo.CreateUser(user)
}


func (u *authUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
