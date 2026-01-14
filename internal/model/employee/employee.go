package employee

import (
	"os/user"
	"time"
)

type EmployeeStatus string

const (
	StatusActive   EmployeeStatus = "active"
	StatusInActive EmployeeStatus = "inactive"
)

type Employee struct {
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint
	User   user.User `gorm:"constraint:onDelete:CASCADE"`

	Name      string    `json:"name"`
	Phone     string    `json:"phone" gorm:"unique"`
	Address   string    `json:"address"`
	Status    string    `json:"status" gorm:"type:enum('active','inactive');default:'active"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
