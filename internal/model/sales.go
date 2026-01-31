package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type SalesStatus string

const (
	StatusUnpaid    SalesStatus = "unpaid"
	StatusPartial   SalesStatus = "partial"
	StatusPaid      SalesStatus = "paid"
	StatusCancelled SalesStatus = "cancelled"
)

type Sales struct {
	ID         uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID uint     `json:"customer_id" gorm:"not null"`
	Customer   Customer `json:"customer" gorm:"foreignKey:CustomerID"`

	SalesDate       time.Time       `json:"sales_date"`
	TotalAmount     decimal.Decimal `json:"total_amount"`
	PaidAmount      decimal.Decimal `json:"paid_amount"`
	RemainingAmount decimal.Decimal `json:"remaining_amount"`
	Status          SalesStatus     `json:"status"`
	Note            string          `json:"note"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}
