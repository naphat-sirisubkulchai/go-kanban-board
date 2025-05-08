package models

import "time"

type Task struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title      string    `gorm:"not null"`
	Description string
	ColumnID   string    `gorm:"not null"`
	Position   int       `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Column Column `gorm:"foreignKey:ColumnID"`
	Tags   []Tag  `gorm:"many2many:task_tags;"`
	Assignees []User `gorm:"many2many:task_assignees;"`
}
