package model

import (
	"gorm.io/gorm"
)

type EmployeeStatus string

const (
	StatusActive   EmployeeStatus = "active"
	StatusInActive EmployeeStatus = "inactive"
)

type Employee struct {
	gorm.Model
	UserID uint `json:"user_id" gorm:"uniqueIndex;not null"`
	User   User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`

	Name    string         `json:"name"`
	Phone   string         `json:"phone" gorm:"unique"`
	Address string         `json:"address"`
	Status  EmployeeStatus `json:"status" gorm:"type:enum('active','inactive');default:'active"`
	Photo   string         `json:"photo"`
}
