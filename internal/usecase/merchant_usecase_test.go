package usecase

import (
	"testing"

	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/mock/repo_mock"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	merchantRepo := new(repo_mock.MerchantRepoMock)
	useCase := NewMerchantUseCase(merchantRepo)

	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	merchantRepo.On("Create", merchant).Return(merchant, nil)

	result, err := useCase.RegisterNewMerchant(merchant)
	assert.NoError(t, err)
	assert.Equal(t, merchant.IdMerchant, result.IdMerchant)

	merchantRepo.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	mockRepo := new(repo_mock.MerchantRepoMock)
	useCase := NewMerchantUseCase(mockRepo)

	merchants := []entity.Merchant{
		{
			IdMerchant:   "uuid-merchant-test",
			IdUser:       "uuid-user-test",
			NameMerchant: "name-merchant-test",
			Address:      "address-test",
			IdProduct:    "uuid-product-test",
			Balance:      10000,
		},
		{
			IdMerchant:   "uuid-merchant-test-2",
			IdUser:       "uuid-user-test-2",
			NameMerchant: "name-merchant-test-2",
			Address:      "address-test-2",
			IdProduct:    "uuid-product-test-2",
			Balance:      20000,
		},
	}

	mockRepo.On("List").Return(merchants, nil)

	result, err := useCase.FindAllMerchant()
	assert.NoError(t, err)
	assert.Len(t, result, len(merchants))

	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(repo_mock.MerchantRepoMock)
	useCase := NewMerchantUseCase(mockRepo)

	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	mockRepo.On("Get", "1").Return(merchant, nil)

	result, err := useCase.FindMerchantByID("1")
	assert.NoError(t, err)
	assert.Equal(t, merchant, result)

	mockRepo.AssertExpectations(t)
}
