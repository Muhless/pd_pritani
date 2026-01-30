package model

import "time"

type Sales struct {
	ID        uint `json:"id" gorm:"primaryKey;autoIncrement"`
	SalesDate time.Time 
}
