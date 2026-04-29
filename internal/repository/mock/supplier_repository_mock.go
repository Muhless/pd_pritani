package mock

import (
	"pd_pritani/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockSupplierRepository struct {
	mock.Mock
}

func (m *MockSupplierRepository) FindAll(page, limit int) ([]model.Supplier, int64, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]model.Supplier), args.Get(1).(int64), args.Error(2)
}

func (m *MockSupplierRepository) FindByID(id uint) (*model.Supplier, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Supplier), args.Error(1)
}

func (m *MockSupplierRepository) Create(supplier *model.Supplier) error {
	args := m.Called(supplier)
	return args.Error(0)
}

func (m *MockSupplierRepository) Update(supplier *model.Supplier) error {
	args := m.Called(supplier)
	return args.Error(0)
}

func (m *MockSupplierRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

