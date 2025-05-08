package models

import "time"

type Tag struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	Color     string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Tasks []Task `gorm:"many2many:task_tags;"`
}
