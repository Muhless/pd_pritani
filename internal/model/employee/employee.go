package employee

import (
	"time"
)

type EmployeeStatus string

const (
	StatusActive   EmployeeStatus = "active"
	StatusInActive EmployeeStatus = "inactive"
)

type Employee struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone" gorm:"unique"`
	Address   string    `json:"address"`
	Photo     string    `json:"photo"`
	Status    string    `json:"status" gorm:"type:enum('active','inactive');default:'active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
