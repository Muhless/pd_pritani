package repository

import (
	"pd_pritani/internal/model"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SalesRepository interface {
	FindAll() ([]model.Sales, error)
	FindById(id uint) (*model.Sales, error)
	Create(sales *model.Sales) error
	Update(sales *model.Sales) error
	Delete(id uint) error

	GetTotalRevenue(month, year int) (decimal.Decimal, error)
	GetRecentSales(limit int) ([]model.Sales, error)
}

type salesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) SalesRepository {
	return &salesRepository{db}
}

func (r *salesRepository) FindAll() ([]model.Sales, error) {
	var sales []model.Sales
	err := r.db.Preload("Customer").Preload("Employee").Preload("SalesItems.Product").Find(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *salesRepository) FindById(id uint) (*model.Sales, error) {
	var sales model.Sales
	err := r.db.Preload("Customer").Preload("Employee").Preload("SalesItems.Product").Find(&sales, id).Error

	if err != nil {
		return nil, err
	}
	return &sales, nil
}

func (r *salesRepository) Create(sales *model.Sales) error {
	return r.db.Create(sales).Error
}

func (r *salesRepository) Update(sales *model.Sales) error {
	return r.db.Save(sales).Error
}

func (r *salesRepository) Delete(id uint) error {
	return r.db.Delete(&model.Sales{}, id).Error
}

func (r *salesRepository) GetTotalRevenue(month, year int) (decimal.Decimal, error) {
	var total decimal.Decimal

	err := r.db.Model(&model.Sales{}).
		Select("COALESCE(SUM(total_price),0)").
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT (YEAR FROM created_at)=?", month, year).
		Where("status = ? ", model.SalesStatusPaid).
		Where("deleted_at IS NULL").
		Scan(&total).Error
	return total, err
}

func (r *salesRepository) GetRecentSales(limit int) ([]model.Sales, error) {
	var sales []model.Sales
	err := r.db.Preload("Customer").Preload("Employee").
		Order("created_at DESC").
		Limit(limit).
		Find(&sales).Error
	return sales, err
}
