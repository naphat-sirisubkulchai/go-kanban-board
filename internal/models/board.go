package models

import "time"

type Board struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	OwnerID   string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Owner User `gorm:"foreignKey:OwnerID"`
	Members []User `gorm:"many2many:board_members"`

}
