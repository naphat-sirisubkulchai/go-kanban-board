package main

import (
	"log"

	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/config"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/handler"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/infrastucture"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/middleware"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/repository"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load config
	config.LoadEnv()

	// Connect external dependencies
	db := infrastructure.ConnectDatabase()
	infrastructure.ConnectRedis()

	// Create app
	app := fiber.New()

	// Repositories
	authRepo := repository.NewAuthRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)
	userRepo := repository.NewUserRepository(db)
	boardRepo := repository.NewBoardRepository(db)
	columnRepo := repository.NewColumnRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Usecases
	authUsecase := usecase.NewAuthUsecase(authRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	boardUsecase := usecase.NewBoardUsecase(boardRepo, notificationRepo)
	columnUsecase := usecase.NewColumnUsecase(columnRepo)
	taskUsecase := usecase.NewTaskUsecase(taskRepo, notificationRepo)
	notificationUsecase := usecase.NewNotificationUsecase(notificationRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	boardHandler := handler.NewBoardHandler(boardUsecase)
	columnHandler := handler.NewColumnHandler(columnUsecase)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	notificationHandler := handler.NewNotificationHandler(notificationUsecase)





	api := app.Group("/api")
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)


	api.Get("/users", middleware.JWTProtected(), userHandler.GetAllUsers)
	api.Get("/users/:email", middleware.JWTProtected(), userHandler.GetUserByEmail)

	board := api.Group("/boards", middleware.JWTProtected())
	board.Post("/", boardHandler.CreateBoard)
	board.Put("/:id", boardHandler.RenameBoard)
	board.Delete("/:id", boardHandler.DeleteBoard)
	board.Post("/:id/invite", boardHandler.InviteMember)
	board.Get("/name/:name", boardHandler.GetBoardByName)

	column := api.Group("/columns",middleware.JWTProtected())
	column.Post("/", columnHandler.CreateColumn)
	column.Patch("/:id", columnHandler.UpdateColumnName)
	column.Delete("/:id", columnHandler.DeleteColumn)
	column.Get("/board/:board_id", columnHandler.GetColumnsByBoard)

	task := api.Group("/tasks",middleware.JWTProtected())
	task.Post("/", taskHandler.CreateTask)
	task.Patch("/:id", taskHandler.UpdateTask)
	task.Delete("/:id", taskHandler.DeleteTask)
	task.Get("/column/:column_id", taskHandler.GetTasksByColumn)
	task.Patch("/reorder/:id", taskHandler.ReorderTask)
	task.Post("/:task_id/tag/:tag_id", taskHandler.AddTagToTask)
	task.Delete("/:task_id/tag/:tag_id", taskHandler.RemoveTagFromTask)
	task.Post("/:task_id/assign/:user_id", taskHandler.AssignUserToTask)

	notif := api.Group("/notifications",middleware.JWTProtected())
	notif.Get("/:user_id", notificationHandler.GetNotifications)
	notif.Patch("/:notification_id/read", notificationHandler.MarkAsRead)


	log.Fatal(app.Listen(":3000"))
}

