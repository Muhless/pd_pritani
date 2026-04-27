package service_test

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
