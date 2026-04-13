package service

import (
	"errors"
	"fmt"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SalesItemRequest struct {
	ProductID uint            `json:"product_id" binding:"required"`
	Quantity  decimal.Decimal `json:"quantity" binding:"required"`
}

type SalesRequest struct {
	CustomerID uint               `json:"customer_id" binding:"required"`
	Notes      string             `json:"notes"`
	Items      []SalesItemRequest `json:"items" binding:"required,min=1"`
}

type SalesService interface {
	GetAll() ([]model.Sales, error)
	GetByID(id uint) (*model.Sales, error)
	Create(employeeID uint, req SalesRequest) error
	UpdateStatus(id uint, status string) error
	Delete(id uint) error
}

type salesService struct {
	db           *gorm.DB
	salesRepo    repository.SalesRepository
	productRepo  repository.ProductRepository
	EmployeeRepo repository.EmployeeRepository
}

func NewSalesService(
	db *gorm.DB,
	salesRepo repository.SalesRepository,
	productRepo repository.ProductRepository,
	employeeRepo repository.EmployeeRepository,
) SalesService {
	return &salesService{db, salesRepo, productRepo, employeeRepo}
}

func generateInvoiceNumber(id uint) string {
	now := time.Now()
	return fmt.Sprintf("INV-%d%02d%02d-%04d", now.Year(), now.Month(), now.Day(), id)
}

func (s *salesService) GetAll() ([]model.Sales, error) {
	sales, err := s.salesRepo.FindAll()
	if err != nil {
		return nil, errors.New("fetching data failed")
	}
	return sales, nil
}

func (s *salesService) GetByID(id uint) (*model.Sales, error) {
	sales, err := s.salesRepo.FindById(id)
	if err != nil {
		return nil, errors.New("sales not found")
	}
	return sales, nil
}

func (s *salesService) Create(userID uint, req SalesRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		employee, err := s.EmployeeRepo.FindByUserId(userID)
		if err != nil {
			return errors.New("employee not found")
		}

		var totalPrice decimal.Decimal
		var salesItems []model.SalesItems

		for _, item := range req.Items {
			// product
			product, err := s.productRepo.FindByID(item.ProductID)
			if err != nil {
				return fmt.Errorf("product not found: %d", item.ProductID)
			}
			// stock
			if product.Stock.LessThan(item.Quantity) {
				return fmt.Errorf("product stock limit: %s", product.Name)
			}

			// subtotal
			subtotal := product.Price.Mul(item.Quantity)
			totalPrice = totalPrice.Add(subtotal)

			salesItems = append(salesItems, model.SalesItems{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price,
				Subtotal:  subtotal,
			})

			// deprecete stock
			product.Stock = product.Stock.Sub(item.Quantity)
			if err := tx.Save(product).Error; err != nil {
				return errors.New("update stock failed")
			}
		}

		sales := &model.Sales{
			CustomerID: req.CustomerID,
			EmployeeID: employee.ID,
			TotalPrice: totalPrice,
			Status:     model.SalesStatusPending,
			Notes:      req.Notes,
			SalesItems: salesItems,
		}
		if err := tx.Create(sales).Error; err != nil {
			return errors.New("failed creating sales")
		}

		sales.InvoiceNumber = generateInvoiceNumber(sales.ID)
		if err := tx.Save(sales).Error; err != nil {
			return errors.New("update invoice failed")
		}
		return nil
	})
}

func (s *salesService) UpdateStatus(id uint, status string) error {
	sales, err := s.salesRepo.FindById(id)
	if err != nil {
		return errors.New("Sales not found")
	}
	sales.Status = model.SalesStatus(status)
	return s.salesRepo.Update(sales)
}

func (s *salesService) Delete(id uint) error {
	_, err := s.salesRepo.FindById(id)
	if err != nil {
		return errors.New("id not found")
	}
	return s.salesRepo.Delete(id)
}
