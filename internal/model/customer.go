package model

type Customer struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Phone 
	Address string `json:"address"`
}
