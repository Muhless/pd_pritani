package service

import (
	"pd_pritani/internal/dto"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
)

type SupplierService interface {
	GetAll() ([]model.Supplier, error)
	GetByID(id uint) (*model.Supplier, error)
	Create(req dto.CreateSupplierRequest) (*model.Supplier, error)
	Update(id uint, req dto.UpdateSupplierRequest) (*model.Supplier, error)
	Delete(id uint) error
}

type supplierService struct {
	repo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierService{repo}
}

func (s *supplierService) GetAll() ([]model.Supplier, error) {
	return s.repo.FindAll()
}

func (s *supplierService) GetByID(id uint) (*model.Supplier, error) {
	return s.repo.FindByID(id)
}

func (s *supplierService) Create(req dto.CreateSupplierRequest) (*model.Supplier, error) {
	supplier := &model.Supplier{
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
		Notes:   req.Notes,
	}
	if err := s.repo.Create(supplier); err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *supplierService) Update(id uint, req dto.UpdateSupplierRequest) (*model.Supplier, error) {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		supplier.Name = req.Name
	}
	if req.Phone != "" {
		supplier.Phone = req.Phone
	}
	if req.Address != "" {
		supplier.Address = req.Address
	}
	if req.Notes != "" {
		supplier.Notes = req.Notes
	}
	if err := s.repo.Update(supplier); err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *supplierService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}