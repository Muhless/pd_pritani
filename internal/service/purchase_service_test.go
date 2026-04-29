package service_test

import (
	"errors"
	"pd_pritani/internal/dto"
	"pd_pritani/internal/model"
	repoMock "pd_pritani/internal/repository/mock"
	"pd_pritani/internal/service"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllPurchase_Success(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)

	purchases := []model.Purchase{
		{PONumber: "PO-20260429-001", TotalPrice: decimal.NewFromInt(100000)},
		{PONumber: "PO-20260429-002", TotalPrice: decimal.NewFromInt(300000)},
	}

	mockRepo.On("FindAll", 1, 10).Return(purchases, int64(2), nil)
	purchasesService := service.NewPurchaseService(mockRepo)
	result, total, err := purchasesService.GetAll(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, int64(2), total)
	mockRepo.AssertExpectations(t)

}

func TestGetAllPurchase_Failed(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)
	mockRepo.On("FindAll", 1, 10).Return([]model.Purchase{}, int64(0), errors.New("db error"))

	purchaseService := service.NewPurchaseService(mockRepo)
	result, total, err := purchaseService.GetAll(1, 10)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, int64(0), total)
	mockRepo.AssertExpectations(t)
}

func TestPurchaseGetByID_Success(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)

	purchase := &model.Purchase{
		PONumber:   "PO-20260101-001",
		TotalPrice: decimal.NewFromInt(100000),
	}

	mockRepo.On("FindByID", uint(1)).Return(purchase, nil)
	purchaseService := service.NewPurchaseService(mockRepo)
	result, err := purchaseService.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "PO-20260101-001", result.PONumber)
	mockRepo.AssertExpectations(t)

}

func TestGetPurchaseByID_NotFound(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)
	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("id not found"))

	ps := service.NewPurchaseService(mockRepo)
	result, err := ps.GetByID(99)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "id not found", err.Error())
	mockRepo.AssertExpectations(t)

}

func TestCreatePurchase_Success(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)
	purchases := &model.Purchase{
		PONumber:   "PO-123-001",
		EmployeeID: 1,
		SupplierID: 1,
		TotalPrice: decimal.NewFromInt(1000000),
		Status:     model.PurchaseStatusPending,
	}

	mockRepo.On("Create", mock.AnythingOfType("*model.Purchase"), mock.AnythingOfType("[]model.PurchaseItem")).Return(nil)
	mockRepo.On("FindByID", uint(0)).Return(purchases, nil)

	purchaseService := service.NewPurchaseService(mockRepo)
	result, err := purchaseService.Create(1, dto.CreatePurchaseRequest{
		SupplierID: 1,
		Items: []dto.CreatePurchaseItemRequest{
			{
				ProductID: 1,
				Quantity:  decimal.NewFromInt(10),
				Price:     decimal.NewFromInt(1000000),
			},
		},
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)

}

func TestUpdatePurchaseStatus_Success(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)
	purchase := &model.Purchase{
		PONumber: "PO-111-001",
		Status:   model.PurchaseStatusPending,
	}

	mockRepo.On("FindByID", uint(1)).Return(purchase, nil)
	mockRepo.On("UpdateStatus", mock.AnythingOfType("*model.Purchase")).Return(nil)

	servicePurchase := service.NewPurchaseService(mockRepo)
	result, err := servicePurchase.UpdateStatus(1, dto.UpdatePurchaseStatusRequest{
		Status: "received",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)

}

func TestUpdatePurchaseStatus_NotFound(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)

	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("id not found"))

	ps := service.NewPurchaseService(mockRepo)
	result, err := ps.UpdateStatus(99, dto.UpdatePurchaseStatusRequest{
		Status: "received",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)

}

func TestDeletePurchase_Success(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)
	purchase := &model.Purchase{
		PONumber: "PO-1234-001",
	}

	mockRepo.On("FindByID", uint(1)).Return(purchase, nil)
	mockRepo.On("Delete", uint(1)).Return(nil)

	ps := service.NewPurchaseService(mockRepo)
	err := ps.Delete(uint(1))

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeletePurchase_NotFound(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)

	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("id not found"))
	sp := service.NewPurchaseService(mockRepo)
	err := sp.Delete(99)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
