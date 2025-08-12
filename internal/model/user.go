package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
