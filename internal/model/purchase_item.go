package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type PurchaseItem struct {
	gorm.Model
	PurchaseID uint            `json:"purchase_id" gorm:"not null"`
	ProductID  uint            `json:"product_id" gorm:"not null"`
	Quantity   decimal.Decimal `json:"quantity" gorm:"type:numeric(12,2);not null"`
	Price      decimal.Decimal `json:"price" gorm:"type:numeric(12,2);not null"`
	Subtotal   decimal.Decimal `json:"subtotal" gorm:"type:numeric(12,2);not null"`

	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
