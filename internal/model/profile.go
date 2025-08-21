package model

import "time"

type Profile struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"uniqueIndex"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	Phone     string    `json:"phone" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}
