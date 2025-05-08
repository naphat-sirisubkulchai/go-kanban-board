package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/usecase"
)

type ColumnHandler struct {
	usecase usecase.ColumnUsecase
}

func NewColumnHandler(u usecase.ColumnUsecase) *ColumnHandler {
	return &ColumnHandler{usecase: u}
}

func (h *ColumnHandler) CreateColumn(c *fiber.Ctx) error {
	var col models.Column
	if err := c.BodyParser(&col); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.usecase.Create(&col); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(col)
}

func (h *ColumnHandler) UpdateColumnName(c *fiber.Ctx) error {
	columnID := c.Params("id")
	var body struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.usecase.UpdateName(columnID, body.Name); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Column name updated"})
}

func (h *ColumnHandler) DeleteColumn(c *fiber.Ctx) error {
	columnID := c.Params("id")
	if err := h.usecase.Delete(columnID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Column deleted"})
}

func (h *ColumnHandler) GetColumnsByBoard(c *fiber.Ctx) error {
	boardID := c.Params("board_id")
	columns, err := h.usecase.GetByBoardID(boardID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(columns)
}
