package model

import "time"

type SalesItems struct {
	ID      uint  `json:"id" gorm:"primaryKey;autoIncrement"`
	SalesID uint  `json:"sales_id" gorm:"not null"`
	Sales   Sales `json:"sales" gorm:"foreignKey:SalesID"`

	ProductID uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`

	Quantity int `json:"quantity" gorm:"default:0"`
	Subtotal int `json:"subtotal" gorm:"default:0"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
