package service

import (
	"errors"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"

	"github.com/shopspring/decimal"
)

type ProductRequest struct {
	Name        string          `json:"name" binding:"required"`
	Category    string          `json:"category" binding:"required,oneof=rice bran"`
	Stock       decimal.Decimal `json:"stock"`
	Price       decimal.Decimal `json:"price"`
	Unit        string          `json:"unit" binding:"required"`
	Photo       string          `json:"photo"`
	Description string          `json:"description"`
}

type ProductUpdateRequest struct {
	Name        string          `json:"name"`
	Category    string          `json:"category" binding:"omitempty,oneof=rice bran"`
	Stock       decimal.Decimal `json:"stock"`
	Price       decimal.Decimal `json:"price"`
	Unit        string          `json:"unit"`
	Photo       string          `json:"photo"`
	Description string          `json:"description"`
}

type ProductService interface {
	GetAll(page, limit int) ([]model.Product, int64, error)
	GetByID(id uint) (*model.Product, error)
	Create(req ProductRequest) error
	Update(id uint, req ProductUpdateRequest) error
	Delete(id uint) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo}
}

func (s *productService) GetAll(page, limit int) ([]model.Product, int64, error) {
	products, total, err := s.productRepo.FindAll(page, limit)
	if err != nil {
		return nil, 0, errors.New("failed")
	}
	return products, total, nil
}

func (s *productService) GetByID(id uint) (*model.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (s *productService) Create(req ProductRequest) error {
	product := &model.Product{
		Name:        req.Name,
		Category:    model.ProductCategory(req.Category),
		Stock:       req.Stock,
		Price:       req.Price,
		Unit:        req.Unit,
		Photo:       req.Photo,
		Description: req.Description,
	}
	return s.productRepo.Create(product)
}

func (s *productService) Update(id uint, req ProductUpdateRequest) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Category != "" {
		product.Category = model.ProductCategory(req.Category)
	}
	if !req.Stock.IsZero() {
		product.Stock = req.Stock
	}
	if !req.Price.IsZero() {
		product.Price = req.Price
	}
	if req.Unit != "" {
		product.Unit = req.Unit
	}
	if req.Photo != "" {
		product.Photo = req.Photo
	}
	if req.Description != "" {
		product.Description = req.Description
	}

	return s.productRepo.Update(product)
}

func (s *productService) Delete(id uint) error {
	_, err := s.productRepo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}
	return s.productRepo.Delete(id)

}
