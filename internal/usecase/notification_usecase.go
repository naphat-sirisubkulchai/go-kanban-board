package usecase

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/repository"
)

type NotificationUsecase interface {
	GetByUserID(userID string) ([]models.Notification, error)
	MarkAsRead(notificationID string) error

}

type notificationUsecase struct {
	repo repository.NotificationRepository
}

func NewNotificationUsecase(r repository.NotificationRepository) NotificationUsecase {
	return &notificationUsecase{repo: r}
}

func (u *notificationUsecase) GetByUserID(userID string) ([]models.Notification, error) {
	return u.repo.GetNotificationsByUserID(userID)
}
func (u *notificationUsecase) MarkAsRead(notificationID string) error {
	return u.repo.MarkAsRead(notificationID)
}
