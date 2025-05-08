package models

import "time"

type Notification struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string    `gorm:"not null"`
	Message   string    `gorm:"not null"`
	IsRead    bool      `gorm:"default:false"`
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}
