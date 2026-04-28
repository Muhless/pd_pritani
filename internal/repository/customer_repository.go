package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll(page, limit int) ([]model.Customer, int64, error)
	FindByID(id uint) (*model.Customer, error)
	Create(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) FindAll(page,limit int) ([]model.Customer, int64, error) {
	var customers []model.Customer
	var total int64

	offset := (page-1) * limit

	r.db.Model(&model.Customer{}).Count(&total)

	err := r.db.Offset(offset).Limit(limit).Find(&customers).Error
	if err != nil {
		return nil,0, err
	}
	return customers, total, err
}

func (r *customerRepository) FindByID(id uint) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Create(customer *model.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) Update(customer *model.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return r.db.Delete(&model.Customer{}, id).Error
}
