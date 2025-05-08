package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/config"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
)

type NotificationRepository interface {
	CreateNotification(n *models.Notification) error
	GetNotificationsByUserID(userID string) ([]models.Notification, error)
	MarkAsRead(notificationID string) error

}

type notificationRepo struct{}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepo{}
}

func (r *notificationRepo) CreateNotification(n *models.Notification) error {
	return config.DB.Create(n).Error
}
func (r *notificationRepo) GetNotificationsByUserID(userID string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := config.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepo) MarkAsRead(notificationID string) error {
	return config.DB.Model(&models.Notification{}).
		Where("id = ?", notificationID).
		Update("is_read", true).Error
}