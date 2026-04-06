package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductCategory string

const (
	ProductCategoryRice ProductCategory = "rice"
	ProductCategoryMusk ProductCategory = "bran"
)

type Product struct {
	gorm.Model
	Name        string          `json:"name" gorm:"not null"`
	Category    ProductCategory `json:"category" gorm:"type:product_category;default:'rice'"`
	Stock       decimal.Decimal `json:"stock" gorm:"default:0"`
	Unit        string          `json:"unit" gorm:"type:varchar(20); not null"`
	Price       decimal.Decimal `json:"price" gorm:"type:numeric(12,2);not null"`
	Photo       string          `json:"photo"`
	Description string          `json:"description" gorm:"type:text"`
}
