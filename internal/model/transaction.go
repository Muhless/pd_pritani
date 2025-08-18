package model

import "time"

// This is only enum
type TransactionStatus string

const (
	TransactionOnProccess TransactionStatus = "diproses"
	TransactionSuccess    TransactionStatus = "selesai"
	TransactionFailed     TransactionStatus = "dibatalkan"
)

// Status using GORM + Enum
type Transaction struct {
	ID        uint              `json:"id" gorm:"primaryKey;autoIncrement"`
	Status    TransactionStatus `json:"status" gorm:"type:enum('selesai', 'dibatalkan', 'diproses);default:'diproses'"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`

	// Relation to Procuct table
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
