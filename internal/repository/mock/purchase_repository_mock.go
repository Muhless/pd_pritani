package mock

import (
	"pd_pritani/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockPurchaseRepository struct {
	mock.Mock
}

func (m *MockPurchaseRepository) FindAll(page, limit int) ([]model.Purchase, int64, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]model.Purchase), args.Get(1).(int64), args.Error(2)
}

func (m *MockPurchaseRepository) FindByID(id uint) (*model.Purchase, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Purchase), args.Error(1)

}

func (m *MockPurchaseRepository) Create(purchase *model.Purchase) error {
	args := m.Called(purchase)
	return args.Error(0)
}

func (m *MockPurchaseRepository) Update(purchase *model.Purchase) error {
	args := m.Called(purchase)
	return args.Error(0)
}

func (m *MockPurchaseRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
