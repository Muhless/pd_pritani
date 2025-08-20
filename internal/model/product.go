package model

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Photo     string    `json:"photo"`
	Type      string    `json:"type"`
	Stock     int       `json:"stock" gorm:"default:0"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateProductInput struct {
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Stock *int    `json:"stock"`
	Price *int    `json:"price"`
	Photo *string `json:"photo"`
}
