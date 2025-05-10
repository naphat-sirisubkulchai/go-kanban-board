package repository

import (
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(n *models.Notification) error
	GetNotificationsByUserID(userID string) ([]models.Notification, error)
	MarkAsRead(notificationID string) error
}

type notificationRepo struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepo{db: db}
}

func (r *notificationRepo) CreateNotification(n *models.Notification) error {
	return r.db.Create(n).Error
}

func (r *notificationRepo) GetNotificationsByUserID(userID string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepo) MarkAsRead(notificationID string) error {
	return r.db.Model(&models.Notification{}).
		Where("id = ?", notificationID).
		Update("is_read", true).Error
}
