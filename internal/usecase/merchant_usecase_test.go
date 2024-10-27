package usecase

import (
	"errors"
	"testing"

	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/mock/repo_mock"

	"github.com/stretchr/testify/suite"
)

type merchantUsecaseSuite struct {
	suite.Suite
	merchantRepo    *repo_mock.MerchantRepoMock
	merchantUsecase MerchantUseCase
	log             logger.Logger
}

func TestMerchantUsecaseSuite(t *testing.T) {
	suite.Run(t, new(merchantUsecaseSuite))
}

func (m *merchantUsecaseSuite) SetupTest() {
	m.merchantRepo = new(repo_mock.MerchantRepoMock)
	m.log = logger.NewLogger()
	m.merchantUsecase = NewMerchantUseCase(m.merchantRepo, &m.log)
}

func (m *merchantUsecaseSuite) TestCreateMerchant_success() {
	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	m.merchantRepo.On("Create", merchant).Return(merchant, nil)

	result, err := m.merchantUsecase.RegisterNewMerchant(merchant)
	m.NoError(err)
	m.Equal(merchant.IdMerchant, result.IdMerchant)
}

func (m *merchantUsecaseSuite) TestGetAllMerchant_success() {
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

	m.merchantRepo.On("List").Return(merchants, nil)

	result, err := m.merchantUsecase.FindAllMerchant()
	m.NoError(err)
	m.Len(result, len(merchants))
}

func (m *merchantUsecaseSuite) TestGetByIDMerchant_success() {
	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	m.merchantRepo.On("Get", "uuid-merchant-test").Return(merchant, nil)

	result, err := m.merchantUsecase.FindMerchantByID("uuid-merchant-test")
	m.NoError(err)
	m.Equal(merchant, result)
}

func (m *merchantUsecaseSuite) TestUpdateMerchant_success() {
	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	m.merchantRepo.On("Get", merchant.IdMerchant).Return(merchant, nil)
	m.merchantRepo.On("Update", merchant, merchant).Return(merchant, nil)

	result, err := m.merchantUsecase.UpdateMerchant(merchant)
	m.NoError(err)
	m.Equal(merchant.IdMerchant, result.IdMerchant)
}

func (m *merchantUsecaseSuite) TestUpdateMerchant_failed() {
	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	m.merchantRepo.On("Get", merchant.IdMerchant).Return(entity.Merchant{}, errors.New("merchant ID of \\uuid-merchant-test\\ not found"))

	result, err := m.merchantUsecase.UpdateMerchant(merchant)
	m.Error(err)
	m.EqualError(err, "merchant ID of \\uuid-merchant-test\\ not found")
	m.Equal(entity.Merchant{}, result)
}

func (m *merchantUsecaseSuite) TestDeleteMerchant_success() {
	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "name-merchant-test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}

	m.merchantRepo.On("Get", merchant.IdMerchant).Return(merchant, nil)
	m.merchantRepo.On("Delete", merchant.IdMerchant).Return(nil)

	err := m.merchantUsecase.DeleteMerchant(merchant.IdMerchant)
	m.NoError(err)
}

func (m *merchantUsecaseSuite) TestDeleteMerchant_failed() {
	merchant := entity.Merchant{
		IdMerchant: "uuid-merchant-test",
	}

	m.merchantRepo.On("Get", merchant.IdMerchant).Return(entity.Merchant{}, errors.New("merchant not found"))

	err := m.merchantUsecase.DeleteMerchant(merchant.IdMerchant)
	m.Error(err)
	m.EqualError(err, "merchant ID of \\uuid-merchant-test\\ not found")
}
