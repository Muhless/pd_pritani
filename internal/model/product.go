package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID        uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string          `json:"name" gorm:"not null"`
	Photo     string          `json:"photo"`
	Type      string          `json:"type"`
	Stock     int             `json:"stock" gorm:"default:0"`
	Price     decimal.Decimal `json:"price" gorm:"type:numeric(12,2);not null"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type UpdateProductInput struct {
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Stock *int    `json:"stock"`
	Price *int    `json:"price"`
	Photo *string `json:"photo"`
}
