package model

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:varchar(50);not null"`
	Phone   string `json:"phone" gorm:"type:varchar(15);not null"`
	Address string `json:"address" gorm:"type:text"`
	Notes   string `json:"notes" gorm:"type:text"`
}
