package service_test

import (
	"errors"
	"pd_pritani/internal/model"
	repoMock "pd_pritani/internal/repository/mock"
	"pd_pritani/internal/service"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllProduct_Success(t *testing.T) {
	mockRepo := new(repoMock.MockProductRepository)

	products := []model.Product{
		{Name: "Beras Premium", Price: decimal.NewFromInt(12000)},
		{Name: "Beras Bulog", Price: decimal.NewFromInt(8000)},
	}

	mockRepo.On("FindAll", 1, 10).Return(products, int64(2), nil)

	productService := service.NewProductService(mockRepo)
	result, total, err := productService.GetAll(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, int64(2), total)
	mockRepo.AssertExpectations(t)
}

func TestGetAllProduct_Failed(t *testing.T) {
	mockRepo := new(repoMock.MockProductRepository)
	mockRepo.On("FindAll", 1, 10).Return([]model.Product{},int64(0),errors.New("db_error"))

	productService := service.NewProductService(mockRepo)
	result, total, err := productService.GetAll(1, 10)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, int64(0), total)
	assert.Equal(t, "failed", err.Error())
	mockRepo.AssertExpectations(t)

}

// get by ID
func TestGetProductByID_Success(t *testing.T) {
	mockRepo := new(repoMock.MockProductRepository)

	product := &model.Product{
		Name:  "Beras Premium",
		Price: decimal.NewFromInt(12000),
	}

	mockRepo.On("FindByID", uint(1)).Return(product, nil)

	productService := service.NewProductService(mockRepo)
	result, err := productService.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "Beras Premium", result.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetProductByID_NotFound(t *testing.T) {
	mockRepo := new(repoMock.MockProductRepository)

	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("not found"))

	productService := service.NewProductService(mockRepo)
	result, err := productService.GetByID(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "product not found", err.Error())
	mockRepo.AssertExpectations(t)
}

// create
func TestCreateProduct_Success(t *testing.T) {
	mockRepo := new(repoMock.MockProductRepository)
	mockRepo.On("Create", mock.AnythingOfType("*model.Product")).Return(nil)

	productService := service.NewProductService(mockRepo)
	err := productService.Create(service.ProductRequest{
		Name:     "Beras Premium",
		Category: "Rice",
		Stock:    decimal.NewFromInt(10),
		Price:    decimal.NewFromInt(12000),
		Unit:     "kg",
	})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct_NotFound(t *testing.T) {
	mockRepo := new(repoMock.MockProductRepository)
	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("not found"))

	productService := service.NewProductService(mockRepo)
	err := productService.Delete(99)

	assert.Error(t, err)
	assert.Equal(t, "product not found", err.Error())
	mockRepo.AssertExpectations(t)
}
