package service_test

import (
	"errors"
	"pd_pritani/internal/dto"
	"pd_pritani/internal/model"
	repoMock "pd_pritani/internal/repository/mock"
	"pd_pritani/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllSuppliers_Success(t *testing.T) {
	mockRepo := new(repoMock.MockSupplierRepository)

	suppliers := []model.Supplier{
		{Name: "Muhta Nuryadi", Phone: "08871165551"},
		{Name: "Febriyansyah", Phone: "089668582465"},
	}

	mockRepo.On("FindAll", 1, 10).Return(suppliers, int64(2), nil)
	suppliersService := service.NewSupplierService(mockRepo)
	result, total, err := suppliersService.GetAll(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, int64(2), total)
	mockRepo.AssertExpectations(t)
}

func TestGetAllSuppliers_Failed(t *testing.T) {
	mockRepo := new(repoMock.MockSupplierRepository)
	mockRepo.On("FindAll", 1, 10).Return([]model.Supplier{}, int64(0), errors.New("db error"))

	supplierService := service.NewSupplierService(mockRepo)
	result, total, err := supplierService.GetAll(1, 10)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, int64(0), total)
	mockRepo.AssertExpectations(t)
}

func TestGetSupplierByID_Success(t *testing.T) {
	mockRepo := new(repoMock.MockSupplierRepository)

	suppliers := &model.Supplier{
		Name:  "Achmad Supriadi",
		Phone: "08871165551",
	}

	mockRepo.On("FindByID", uint(1)).Return(suppliers, nil)

	suppliersService := service.NewSupplierService(mockRepo)
	result, err := suppliersService.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "08871165551", result.Phone)
	mockRepo.AssertExpectations(t)

}

func TestGetSupplierByID_NotFound(t *testing.T) {
	mockRepo := new(repoMock.MockSupplierRepository)

	mockRepo.On("FindByID", uint(88)).Return(nil, errors.New("id not found"))

	sp := service.NewSupplierService(mockRepo)
	result, err := sp.GetByID(88)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "id not found", err.Error())
	mockRepo.AssertExpectations(t)

}

func TestCreateSupplier_Success(t *testing.T) {
	mockRepo := new(repoMock.MockSupplierRepository)

	mockRepo.On("Create", mock.AnythingOfType("*model.Supplier")).Return(nil)

	supplierService := service.NewSupplierService(mockRepo)
	result, err := supplierService.Create(dto.CreateSupplierRequest{
		Name:  "Muhta Nuryadi",
		Phone: "08871165551",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)

}

func TestUpdateSupplier_Success(t *testing.T) {
	mockRepo := new(repoMock.MockSupplierRepository)
	supplier := &model.Supplier{
		Name:  "Muhta Nuryadi",
		Phone: "08871165551",
	}

	mockRepo.On("FindByID", uint(1)).Return(supplier, nil)
	mockRepo.On("Update", mock.AnythingOfType("*model.Supplier")).Return(nil)

	supplierService := service.NewSupplierService(mockRepo)
	result, err := supplierService.Update(uint(1), dto.UpdateSupplierRequest{
		Name:    "Achmad Supriadi",
		Address: "Kp. Kiara R.06 RW.01",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)

}

func TestDeleteSupplier_Success(t *testing.T) {
	mockRepo := new(repoMock.MockSupplierRepository)

	supplier := &model.Supplier{
		Name:  "Admin",
		Phone: "08871165551",
	}

	mockRepo.On("FindByID", uint(1)).Return(supplier, nil)
	mockRepo.On("Delete", uint(1)).Return(nil)

	supplierService := service.NewSupplierService(mockRepo)
	err := supplierService.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

}

func TestDeleteSupplier_NotFound(t *testing.T) {
	mockRepo := new(repoMock.MockSupplierRepository)
	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("id not found"))
	supplierService := service.NewSupplierService(mockRepo)
	err := supplierService.Delete(99)

	assert.Error(t, err)
	assert.Equal(t, "id not found", err.Error())
	mockRepo.AssertExpectations(t)

}
