package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type SalesRepository interface {
	FindAll() ([]model.Sales, error)
	FindById(id uint) (*model.Sales, error)
	Create(sales *model.Sales) error
	Update(sales *model.Sales) error
	Delete(id uint) error
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
