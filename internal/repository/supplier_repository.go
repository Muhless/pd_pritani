package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	FindAll() ([]model.Supplier, error)
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

func (r *supplierRepository) FindAll() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
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
