package usecase

import (
	"errors"
	"testing"

	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/mock/repo_mock"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	merchantRepo := new(repo_mock.MerchantRepoMock)
	log := logger.NewLogger()
	useCase := NewMerchantUseCase(merchantRepo, &log)

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
	log := logger.NewLogger()
	useCase := NewMerchantUseCase(mockRepo, &log)

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
	log := logger.NewLogger()
	useCase := NewMerchantUseCase(mockRepo, &log)

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

func TestUpdate(t *testing.T) {
	mockRepo := new(repo_mock.MerchantRepoMock)
	log := logger.NewLogger()
	useCase := NewMerchantUseCase(mockRepo, &log)

	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	mockRepo.On("Get", merchant.IdMerchant).Return(merchant, nil)
	mockRepo.On("Update", merchant, merchant).Return(merchant, nil)

	result, err := useCase.UpdateMerchant(merchant)
	assert.NoError(t, err)
	assert.Equal(t, merchant.IdMerchant, result.IdMerchant)

	mockRepo.AssertExpectations(t)
}

func TestUpdateFailed(t *testing.T) {
	mockRepo := new(repo_mock.MerchantRepoMock)
	log := logger.NewLogger()
	useCase := NewMerchantUseCase(mockRepo, &log)

	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	mockRepo.On("Get", merchant.IdMerchant).Return(entity.Merchant{}, errors.New("merchant not found"))

	result, err := useCase.UpdateMerchant(merchant)
	assert.Error(t, err)
	assert.EqualError(t, err, "merchant ID of \\uuid-merchant-test\\ not found")
	assert.Equal(t, entity.Merchant{}, result)

	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(repo_mock.MerchantRepoMock)
	log := logger.NewLogger()
	useCase := NewMerchantUseCase(mockRepo, &log)

	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	mockRepo.On("Get", merchant.IdMerchant).Return(merchant, nil)
	mockRepo.On("Delete", merchant.IdMerchant).Return(nil)

	err := useCase.DeleteMerchant(merchant.IdMerchant)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteFailed(t *testing.T) {
	mockRepo := new(repo_mock.MerchantRepoMock)
	log := logger.NewLogger()
	useCase := NewMerchantUseCase(mockRepo, &log)

	merchant := entity.Merchant{
		IdMerchant: "uuid-merchant-test",
	}

	mockRepo.On("Get", merchant.IdMerchant).Return(entity.Merchant{}, errors.New("merchant not found"))

	err := useCase.DeleteMerchant(merchant.IdMerchant)
	assert.Error(t, err)
	assert.EqualError(t, err, "merchant ID of \\uuid-merchant-test\\ not found")

	mockRepo.AssertExpectations(t)
}
