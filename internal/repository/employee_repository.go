package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *model.Employee) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) Create(employee *model.Employee) error {
	return r.db.Create(employee).Error
}
