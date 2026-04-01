package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *model.Admin) error
	FindByUserID(userID uint) (*model.Admin, error)
	Update(admin *model.Admin) error
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

func (r *adminRepository) FindByUserID(userID uint) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("user_id = ?", userID).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Update(admin *model.Admin) error {
	return r.db.Save(admin).Error
}
