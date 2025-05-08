package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/usecase"
)

type NotificationHandler struct {
	usecase usecase.NotificationUsecase
}

func NewNotificationHandler(u usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{usecase: u}
}

func (h *NotificationHandler) GetNotifications(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	notifications, err := h.usecase.GetByUserID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(notifications)
}
func (h *NotificationHandler) MarkAsRead(c *fiber.Ctx) error {
	notificationID := c.Params("notification_id")
	err := h.usecase.MarkAsRead(notificationID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Notification marked as read"})
}

