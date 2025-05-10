package infrastructure

import (
	"fmt"
	"log"
	"time"

	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/config"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	cfg := config.Env

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("Connected to PostgreSQL")

	if err := autoMigrate(db); err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	return db
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.Column{},
		&models.Task{},
		&models.Tag{},
		&models.Notification{},
	)
}
