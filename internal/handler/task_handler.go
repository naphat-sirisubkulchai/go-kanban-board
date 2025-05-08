package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/usecase"
)

type TaskHandler struct {
	usecase usecase.TaskUsecase
}

func NewTaskHandler(u usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: u}
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.usecase.Create(&task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(task)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	task := &models.Task{
		ID:          id,
		Title:       body.Title,
		Description: body.Description,
	}
	if err := h.usecase.Update(task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Task updated"})
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.usecase.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Task deleted"})
}

func (h *TaskHandler) GetTasksByColumn(c *fiber.Ctx) error {
	columnID := c.Params("column_id")
	tasks, err := h.usecase.GetByColumnID(columnID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tasks)
}

func (h *TaskHandler) ReorderTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		Position int `json:"position"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.usecase.Reorder(id, body.Position); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Position updated"})
}
func (h *TaskHandler) AddTagToTask(c *fiber.Ctx) error {
	taskID := c.Params("task_id")
	tagID := c.Params("tag_id")
	if err := h.usecase.AddTag(taskID, tagID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Tag added to task"})
}

func (h *TaskHandler) RemoveTagFromTask(c *fiber.Ctx) error {
	taskID := c.Params("task_id")
	tagID := c.Params("tag_id")
	if err := h.usecase.RemoveTag(taskID, tagID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Tag removed from task"})
}
func (h *TaskHandler) AssignUserToTask(c *fiber.Ctx) error {
	taskID := c.Params("task_id")
	userID := c.Params("user_id")
	if err := h.usecase.AssignUser(taskID, userID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "User assigned and notified"})
}
