package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type PurchaseRepository interface {
	FindAll() ([]model.Purchase, error)
	FindByID(id uint) (*model.Purchase, error)
	Create(purchase *model.Purchase, items []model.PurchaseItems) error
	UpdateStatus(purchase *model.Purchase) error
	Delete(id uint) error
}

type purchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{db}
}

func (r *purchaseRepository) FindAll() ([]model.Purchase, error) {
	var purchases []model.Purchase
	err := r.db.Find(purchases).Error
	return purchases, err
}

func (r *purchaseRepository) FindByID(id uint) (*model.Purchase, error) {
	var purchase model.Purchase
	err := r.db.Preload("Supplier").Preload("Employee").Preload("PurchaseItems.Product").First(&purchase, id).Error
	return &purchase, err
}

func (r *purchaseRepository) Create(purchase *model.Purchase, items []model.PurchaseItems) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(purchase).Error; err != nil {
			return err
		}
		for i := range items {
			items[i].PurchaseID = purchase.ID
		}
		return tx.Create(&items).Error
	})
}

func (r *purchaseRepository) UpdateStatus(purchase *model.Purchase) error {
	return r.db.Model(purchase).Update("status", purchase.Status).Error
}

func (r *purchaseRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("purchase_id=?", id).Delete(&model.PurchaseItems{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Purchase{}, id).Error
	})
}
