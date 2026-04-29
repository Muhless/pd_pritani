package service_test

import (
	"pd_pritani/internal/model"
	repoMock "pd_pritani/internal/repository/mock"
	"testing"
)

func TestGetAllPurchase_Success(t *testing.T) {
	mockRepo := new(repoMock.MockPurchaseRepository)

	purchases := []model.Purchase{}

}
