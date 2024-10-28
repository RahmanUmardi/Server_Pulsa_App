package repo_mock

import (
	"server-pulsa-app/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MerchantRepoMock struct {
	mock.Mock
}

func (m *MerchantRepoMock) CheckBalanceMerchant(id string) (entity.Merchant, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Merchant), args.Error(1)
}

func (m *MerchantRepoMock) Create(payload entity.Merchant) (entity.Merchant, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Merchant), args.Error(1)
}

func (m *MerchantRepoMock) List() ([]entity.Merchant, error) {
	args := m.Called()
	return args.Get(0).([]entity.Merchant), args.Error(1)
}

func (m *MerchantRepoMock) Get(id string) (entity.Merchant, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Merchant), args.Error(1)
}

func (m *MerchantRepoMock) Update(merchant, newMerchant entity.Merchant) (entity.Merchant, error) {
	args := m.Called(merchant, newMerchant)
	return args.Get(0).(entity.Merchant), args.Error(1)
}

func (m *MerchantRepoMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
