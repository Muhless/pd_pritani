package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SalesItems struct {
	gorm.Model
	SalesID   uint            `json:"sales_id" gorm:"not null"`
	ProductID uint            `json:"product_id" gorm:"not null"`
	Quantity  decimal.Decimal `json:"quantity" gorm:"type:numeric(12,2);not null"`
	Price     decimal.Decimal `json:"price" gorm:"type:numeric:(12,2);not null"`
	Subtotal  decimal.Decimal `json:"subtotal" gorm:"type:numeric(12,2);not null"`

	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
