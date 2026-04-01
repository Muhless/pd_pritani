package repository

import (
	"pd_pritani/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	FindById(id uint) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error

}

func (r *userRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Admin").Preload("Employee").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error

}
