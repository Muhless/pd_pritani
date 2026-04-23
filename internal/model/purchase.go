package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PurchaseStatus string

const (
	PurchaseStatusPending   PurchaseStatus = "pending"
	PurchaseStatusPaid      PurchaseStatus = "paid"
	PurchaseStatusCancelled PurchaseStatus = "cancelled"
)

type Purchase struct {
	gorm.Model
	PONumber   string          `json:"po_number" gorm:"type:varchar(50);not null;unique"`
	EmployeeID uint            `json:"admin_id" gorm:"not null"`
	SupplierID uint            `json:"supplier_id" gorm:"not null"`
	TotalPrice decimal.Decimal `json:"total_price" gorm:"type:numeric(12,2);not null"`
	Status     PurchaseStatus  `json:"status" gorm:"varchar(20);not null;default:'pending'"`
	Notes      string          `json:"notes" gorm:"type:text"`

	Admin         *Admin          `json:"admin,omitempty" gorm:"foreignKey:AdminID"`
	Supplier      *Supplier       `json:"supplier,omitempty" gorm:"foreignKey:supplierID"`
	PurchaseItems []PurchaseItems `json:"purchase_items" gorm:"foreignKey:PurchaseID"`
}
