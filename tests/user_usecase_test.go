package usecase_test

import (
	"testing"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *mockUserRepo) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func TestGetUserByEmail(t *testing.T) {
	mockRepo := new(mockUserRepo)
	uc := usecase.NewUserUsecase(mockRepo)

	expected := &models.User{Email: "test@example.com", Name: "Test"}
	mockRepo.On("GetUserByEmail", "test@example.com").Return(expected, nil)

	user, err := uc.GetUserByEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mockUserRepo)
	uc := usecase.NewUserUsecase(mockRepo)

	expectedUsers := []models.User{
		{Email: "user1@example.com", Name: "User1"},
		{Email: "user2@example.com", Name: "User2"},
	}
	mockRepo.On("GetAllUsers").Return(expectedUsers, nil)

	users, err := uc.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}

