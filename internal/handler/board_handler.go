package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/usecase"
)

type BoardHandler struct {
	usecase usecase.BoardUsecase
}

func NewBoardHandler(u usecase.BoardUsecase) *BoardHandler {
	return &BoardHandler{usecase: u}
}

func (h *BoardHandler) CreateBoard(c *fiber.Ctx) error {
	type req struct {
		Title string `json:"title"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	userID := c.Locals("user_id").(string)

	if err := h.usecase.CreateBoard(userID, body.Title); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Board created"})
}

func (h *BoardHandler) RenameBoard(c *fiber.Ctx) error {
	type req struct {
		Title string `json:"title"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	boardID := c.Params("id")
	if err := h.usecase.UpdateBoardTitle(boardID, body.Title); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Board title updated"})
}

func (h *BoardHandler) DeleteBoard(c *fiber.Ctx) error {
	boardID := c.Params("id")

	if err := h.usecase.DeleteBoard(boardID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Board deleted"})
}

func (h *BoardHandler) InviteMember(c *fiber.Ctx) error {
	type req struct {
		UserID string `json:"user_id"`
	}

	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	boardID := c.Params("id")

	if err := h.usecase.InviteMember(boardID, body.UserID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Member invited and notified"})
}
func (h *BoardHandler) GetBoardByName(c *fiber.Ctx) error {
	name := c.Params("name")

	board, err := h.usecase.GetBoardByName(name)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Board not found"})
	}

	return c.JSON(board)
}

