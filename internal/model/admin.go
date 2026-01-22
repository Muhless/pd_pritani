package model

type Admin struct {
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint `json:"user_id" gorm:"uniqueIndex;not null"`
	User   User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`

	Name   string
	Email  string
	Phone  string
	Photo  string
}
