package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:",inline"`
	Role     string `json:"role" gorm:"type:varchar(20);not null"`
	Username string `json:"username" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`

	Admin    *Admin    `json:"admin,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Employee *Employee `json:"employee,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
