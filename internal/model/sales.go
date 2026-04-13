package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SalesStatus string

const (
	SalesStatusPending   SalesStatus = "pending"
	SalesStatusPaid      SalesStatus = "paid"
	SalesStatusCancelled SalesStatus = "cancelled"
)

type Sales struct {
	gorm.Model
	InvoiceNumber string          `json:"invoice_number" gorm:"type:varchar(50);not null;unique"`
	EmployeeID    uint            `json:"employee_id" gorm:"not null"`
	CustomerID    uint            `json:"customer_id" gorm:"not null"`
	TotalPrice    decimal.Decimal `json:"total_price" gorm:"type:numeric(12,2); not null"`
	Status        SalesStatus     `json:"status" gorm:"type:varchar(20);not null;default:'pending'"`
	Notes         string          `json:"notes" gorm:"type:text"`

	Employee   *Employee    `json:"employee" gorm:"foreignKey:EmployeeID"`
	Customer   *Customer    `json:"customer,omitempty" gorm:"foreignKey:CustomerID"`
	SalesItems []SalesItems `json:"sales_items" gorm:"foreignKey:SalesID"`
}
