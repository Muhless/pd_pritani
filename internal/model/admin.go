package model

type Admin struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint
	User   User `gorm:"constraint:onDelete:CASCADE"`

	Name   string
	Email  string
	Phone  string
	Photo  string
	Status string `json:"status" gorm:"type:enum('active','inactive');default:'active"`
}
