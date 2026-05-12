package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(page, limit int) ([]model.Product, int64, error)
	FindByID(id uint) (*model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id uint) error
	GetLowStock(limit int) ([]model.Product, error)
	Count() (int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll(page, limit int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	offset := (page - 1) * limit

	// count total data
	r.db.Model(&model.Product{}).Count(&total)
	// get data with pagination
	err := r.db.Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *productRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

func (r *productRepository) GetLowStock(limit int) ([]model.Product, error) {
	var products []model.Product

	err := r.db.Where("stock < ?", 10).
		Order("stock ASC").
		Limit(limit).
		Find(&products).Error
	return products, err

}

func (r *productRepository) Count() (int64, error) {
	var count int64

	err := r.db.Model(&model.Product{}).Count(&count).Error
	return count, err
}
