package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string `json:"name"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email" gorm:"unique"`
	Phone       string `json:"phone" gorm:"unique"`
	Address     string `json:"address"`

	Sales []Sales `json:"sales,omitempty" gorm:"foreignKey:CustomerID"`
}
