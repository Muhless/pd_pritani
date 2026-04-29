package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	FindAll(page, limit int) ([]model.Supplier, int64, error)
	FindByID(id uint) (*model.Supplier, error)
	Create(supplier *model.Supplier) error
	Update(supplier *model.Supplier) error
	Delete(id uint) error
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db}
}

func (r *supplierRepository) FindAll(page, limit int) ([]model.Supplier, int64, error) {
	var suppliers []model.Supplier
	var total int64

	offset := (page - 1) * limit
	r.db.Model(&model.Supplier{}).Count(&total)

	err := r.db.Offset(offset).Limit(limit).Find(&suppliers).Error
	if err != nil {
		return nil, 0, err

	}
	return suppliers, total, err
}

func (r *supplierRepository) FindByID(id uint) (*model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.First(&supplier, id).Error
	return &supplier, err
}

func (r *supplierRepository) Create(supplier *model.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepository) Update(supplier *model.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *supplierRepository) Delete(id uint) error {
	return r.db.Delete(&model.Supplier{}, id).Error
}
