package models

import "time"

type Column struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	BoardID   string    `gorm:"not null"`
	Position  int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Board Board `gorm:"foreignKey:BoardID"`
}
