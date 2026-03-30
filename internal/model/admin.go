package model

import "gorm.io/gorm"

type AdminStatus string

const (
	AdminStatusActive   AdminStatus = "active"
	AdminStatusInActive AdminStatus = "inactive"
)

type Admin struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"uniqueIndex;not null"`
	User        User   `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Permissions string `json:"permissions" gorm:"not null"`

	Name    string      `json:"name"`
	Email   string      `json:"email" gorm:"unique"`
	Phone   string      `json:"phone" gorm:"unique"`
	Address string      `json:"address"`
	Status  AdminStatus `json:"status" gorm:"type:enum('active','inactive');default:'active'"`
	Photo   string      `json:"photo"`
}
