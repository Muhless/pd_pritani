package model

import (
	"time"
)

type EmployeeStatus string

const (
	StatusActive   EmployeeStatus = "active"
	StatusInActive EmployeeStatus = "inactive"
)

type Employee struct {
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint `json:"user_id" gorm:"uniqueIndex;not null"`
	User   User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`

	Name      string    `json:"name"`
	Phone     string    `json:"phone" gorm:"unique"`
	Address   string    `json:"address"`
	Status    string    `json:"status" gorm:"type:enum('active','inactive');default:'active"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
