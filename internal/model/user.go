package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Role      string    `json:"role" gorm:"type:varchar(20);not null"`
	Username  string    `json:"username" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Admin    *Admin    `json:"admin,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Employee *Employee `json:"employee,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "users"
}
