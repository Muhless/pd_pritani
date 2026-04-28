package service

import (
	"errors"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
	"strings"
)

type CustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone" binding:"required"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

type CustomerService interface {
	GetAll(page, limit int) ([]model.Customer, int64, error)
	GetByID(id uint) (*model.Customer, error)
	Create(req CustomerRequest) error
	Update(id uint, req CustomerRequest) error
	Delete(id uint) error
}

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{customerRepo}
}

func (s *customerService) GetAll(page, limit int) ([]model.Customer, int64, error) {
	customers, total, err := s.customerRepo.FindAll(page, limit)

	if err != nil {
		return nil, 0, errors.New("failed getting customer data")
	}
	return customers, total, nil
}

func (s *customerService) GetByID(id uint) (*model.Customer, error) {
	customer, err := s.customerRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("customer not found")
	}
	return customer, nil

}

func (s *customerService) Create(req CustomerRequest) error {
	customer := &model.Customer{
		Name:        req.Name,
		CompanyName: req.CompanyName,
		Email:       req.Email,
		Phone:       req.Phone,
		Address:     req.Address,
	}

	err := s.customerRepo.Create(customer)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("phone number already used")
		}
		return errors.New("failed")
	}
	return nil
}

func (s *customerService) Update(id uint, req CustomerRequest) error {
	customer, err := s.customerRepo.FindByID(id)
	if err != nil {
		return errors.New("customer not found")
	}
	customer.Name = req.Name
	customer.CompanyName = req.CompanyName
	customer.Email = req.Email
	customer.Phone = req.Phone
	customer.Address = req.Address

	return s.customerRepo.Update(customer)
}

func (s *customerService) Delete(id uint) error {
	_, err := s.customerRepo.FindByID(id)
	if err != nil {
		return errors.New("id not found")
	}
	return s.customerRepo.Delete(id)
}
