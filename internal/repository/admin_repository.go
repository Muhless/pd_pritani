package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *model.Admin) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) Create(admin *model.Admin) error {
	return r.db.Create(admin).Error
}
