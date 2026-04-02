package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *model.Employee) error
	FindByUserId(userID uint) (*model.Employee, error)
	Update(employee *model.Employee) error
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

func (r *employeeRepository) FindByUserId(userID uint) (*model.Employee, error) {
	var employee model.Employee
	err := r.db.Where("user_id = ?", userID).First(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, err
}

func (r *employeeRepository) Update(employee *model.Employee) error {
	return r.db.Save(employee).Error
}
