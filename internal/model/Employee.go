package model

import (
	"time"
)

type Employee struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone" gorm:"unique"`
	Address   string    `json:"address"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}
