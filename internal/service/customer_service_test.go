package service_test

import (
	"errors"
	"pd_pritani/internal/model"
	repoMock "pd_pritani/internal/repository/mock"
	"pd_pritani/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllCustomer_Success(t *testing.T) {
	mockRepo := new(repoMock.MockCustomerRepository)

	customers := []model.Customer{
		{Name: "Muhta Nuryadi", CompanyName: "Toko Maju Jaya", Phone: "08871165551"},
		{Name: "Cihuy", CompanyName: "Toko cihuy abadi", Phone: "089668582565"},
	}

	mockRepo.On("FindAll", 1, 10).Return(customers, int64(2), nil)
	customerService := service.NewCustomerService(mockRepo)
	result, total, err := customerService.GetAll(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, int64(2), total)
	mockRepo.AssertExpectations(t)
}

func TestGetAllCustomer_Failed(t *testing.T) {
	mockRepo := new(repoMock.MockCustomerRepository)

	mockRepo.On("FindAll", 1, 10).Return([]model.Customer{}, int64(2), errors.New("db_error"))

	customerService := service.NewCustomerService(mockRepo)
	result, total, err := customerService.GetAll(1, 10)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, int64(0), total)
	mockRepo.AssertExpectations(t)
}

func TestGetCustomerByID_Success(t *testing.T) {
	mockRepo := new(repoMock.MockCustomerRepository)
	customer := &model.Customer{
		Name:  "Muhta Nuryadi",
		Phone: "08871165551",
	}
	mockRepo.On("FindByID", uint(1)).Return(customer, nil)
	cs := service.NewCustomerService(mockRepo)
	result, err := cs.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Muhta Nuryadi", result.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetCustomerByID_NOtFound(t *testing.T) {
	mockRepo := new(repoMock.MockCustomerRepository)
	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("not found"))

	customerService := service.NewCustomerService(mockRepo)
	result, err := customerService.GetByID(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "customer not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestCreateCustomer_Success(t *testing.T) {
	mockRepo := new(repoMock.MockCustomerRepository)
	mockRepo.On("Create", mock.AnythingOfType("*model.Customer")).Return(nil)

	customerService := service.NewCustomerService(mockRepo)
	err := customerService.Create(service.CustomerRequest{
		Name:        "Muhta Nuryadi",
		CompanyName: "PG. Pritani",
		Phone:       "08871165551",
		Address:     "Kp. Kiara RT.06 RW.01",
	})
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCustomer_Success(t *testing.T) {
	mockRepo := new(repoMock.MockCustomerRepository)
	customer := &model.Customer{
		Name:  "Cihuy",
		Phone: "089668582465",
	}
	mockRepo.On("FindByID", uint(1)).Return(customer, nil)
	mockRepo.On("Update", mock.AnythingOfType(("*model.Customer"))).Return(nil)

	cs := service.NewCustomerService(mockRepo)
	err := cs.Update(1, service.CustomerRequest{
		Name:    "Cihuy Pro Max",
		Phone:   "088118829",
		Address: "kp.cihuy no.96",
	})
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCustomer_success(t *testing.T) {
	mockRepo := new(repoMock.MockCustomerRepository)

	customer := &model.Customer{
		Name:  "cihuy",
		Phone: "08871165551",
	}
	mockRepo.On("FindByID", uint(1)).Return(customer, nil)
	mockRepo.On("Delete", uint(1)).Return(nil)

	cs := service.NewCustomerService(mockRepo)
	err := cs.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCustomer_NotFound(t *testing.T) {
	mockRepo := new(repoMock.MockCustomerRepository)
	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("id not found"))

	cs := service.NewCustomerService(mockRepo)
	err := cs.Delete(99)

	assert.Error(t, err)
	assert.Equal(t, "id not found", err.Error())
	mockRepo.AssertExpectations(t)

}
