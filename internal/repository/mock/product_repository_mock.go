package mock

import (
	"pd_pritani/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindAll(page, limit int) ([]model.Product, int64, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]model.Product), args.Get(1).(int64), args.Error(2)
}

func (m *MockProductRepository) FindByID(id uint) (*model.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductRepository) Create(product *model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Update(product *model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
