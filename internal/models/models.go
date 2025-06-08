package models

import (
	"time"
)

type User struct {
	ID       uint   `json:"id" gorm:"promaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"not null"`
	TeamID   *uint
}

type Tasks struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'new'"`
	Deadline    time.Time `json:"deadline"`
	UserID      uint      `json:"user_id"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
