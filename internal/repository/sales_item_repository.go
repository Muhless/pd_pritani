package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type SalesItemsRepository interface {
	CreateBatch(items []model.SalesItems) error
	DeleteBySalesID(salesID uint) error
}

type salesItemsRepository struct {
	db *gorm.DB
}

func NewSalesItemsRepository(db *gorm.DB) SalesItemsRepository {
	return &salesItemsRepository{db}
}

func (r *salesItemsRepository) CreateBatch(items []model.SalesItems) error {
	return r.db.Create(&items).Error
}

func (r *salesItemsRepository) DeleteBySalesID(SalesID uint) error {
	return r.db.Where("sales_id =?", SalesID).Delete(&model.SalesItems{}).Error
}
