package service

import (
	"errors"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
)

type CustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone" binding:"required"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}

type CustomerService interface {
	GetAll() ([]model.Customer, error)
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

func (s *customerService) GetByID(id uint) (*model.Customer, error) {
	customer, err := s.customerRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("customer nout found")
	}
	return customer, nil

}

func (s *customerService) GetAll() ([]model.Customer, error) {
	customers, err := s.customerRepo.FindAll()
	if err != nil {
		return nil, errors.New("failed getting customer data")
	}
	return customers, nil
}

func (s *customerService) Create(req CustomerRequest) error {
	customer := &model.Customer{
		Name:        req.Name,
		CompanyName: req.CompanyName,
		Email:       req.Email,
		Phone:       req.Phone,
		Address:     req.Address,
	}
	return s.customerRepo.Create(customer)
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
