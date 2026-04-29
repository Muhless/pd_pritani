package service

import (
	"errors"
	"fmt"
	"pd_pritani/internal/dto"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
	"time"

	"github.com/shopspring/decimal"
)

type PurchaseService interface {
	GetAll(page, limit int) ([]model.Purchase, int64, error)
	GetByID(id uint) (*model.Purchase, error)
	Create(employeeID uint, req dto.CreatePurchaseRequest) (*model.Purchase, error)
	UpdateStatus(id uint, req dto.UpdatePurchaseStatusRequest) (*model.Purchase, error)
	Delete(id uint) error
}

type purchaseService struct {
	repo repository.PurchaseRepository
}

func NewPurchaseService(repo repository.PurchaseRepository) PurchaseService {
	return &purchaseService{repo}
}

func (s *purchaseService) GetAll(page, limit int) ([]model.Purchase, int64, error) {
	purchases, total, err := s.repo.FindAll(page, limit)
	if err != nil {
		return nil, 0, errors.New("failed")
	}
	return purchases, total, nil
}

func (s *purchaseService) GetByID(id uint) (*model.Purchase, error) {
	return s.repo.FindByID(id)
}

func (s *purchaseService) Create(employeeID uint, req dto.CreatePurchaseRequest) (*model.Purchase, error) {
	poNumber := fmt.Sprintf("PO-%d-%d", time.Now().Unix(), employeeID)
	total := decimal.Zero
	items := make([]model.PurchaseItem, 0, len(req.Items))

	for _, item := range req.Items {
		subtotal := item.Price.Mul(item.Quantity)
		total = total.Add(subtotal)
		items = append(items, model.PurchaseItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Subtotal:  subtotal,
		})
	}

	purchase := &model.Purchase{
		PONumber:   poNumber,
		EmployeeID: employeeID,
		SupplierID: req.SupplierID,
		TotalPrice: total,
		Status:     model.PurchaseStatusPending,
		Notes:      req.Notes,
	}
	if err := s.repo.Create(purchase, items); err != nil {
		return nil, err
	}
	return s.repo.FindByID(purchase.ID)
}

func (s *purchaseService) UpdateStatus(id uint, req dto.UpdatePurchaseStatusRequest) (*model.Purchase, error) {
	purchase, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	purchase.Status = model.PurchaseStatus(req.Status)
	if err := s.repo.UpdateStatus(purchase); err != nil {
		return nil, err
	}
	return purchase, nil
}

func (s *purchaseService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
