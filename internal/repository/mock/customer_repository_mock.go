package mock

import (
	"pd_pritani/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) FindAll(page, limit int) ([]model.Customer, int64, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]model.Customer), args.Get(1).(int64), args.Error(2)
}

func (m *MockCustomerRepository) FindByID(id uint) (*model.Customer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Create(customer *model.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) Update(customer *model.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
