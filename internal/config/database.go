package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
	"github.com/naphat-sirisubkulchai/go-kanban-board/internal/models"
	"context"
)

var DB *gorm.DB
var RedisClient *redis.Client

func InitDB() {
	_ = godotenv.Load(".env") // optional .env loader

	// เชื่อมต่อ PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DB = db

	
	ConnectRedis()

	
	autoMigrate()
}

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", 
		Password: "",               
		DB:       0,                
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Successfully connected to Redis!")
}

func autoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Board{},
		&models.Column{},
		&models.Task{},
		&models.Tag{},
		&models.Notification{},
	)
	if err != nil {
		log.Fatal("failed auto migration:", err)
	}
}
