package usecase_mock

import (
	"server-pulsa-app/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MerchantUsecaseMock struct {
	mock.Mock
}

func (m *MerchantUsecaseMock) RegisterNewMerchant(payload entity.Merchant) (entity.Merchant, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Merchant), args.Error(1)
}

func (m *MerchantUsecaseMock) FindAllMerchant() ([]entity.Merchant, error) {
	args := m.Called()
	return args.Get(0).([]entity.Merchant), args.Error(1)
}

func (m *MerchantUsecaseMock) FindMerchantByID(id string) (entity.Merchant, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Merchant), args.Error(1)
}

func (m *MerchantUsecaseMock) UpdateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Merchant), args.Error(1)
}

func (m *MerchantUsecaseMock) DeleteMerchant(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
