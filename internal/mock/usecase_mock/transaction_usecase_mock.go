package usecase_mock

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/shared/custom"

	"github.com/stretchr/testify/mock"
)

type MockTransactionUseCase struct {
	mock.Mock
}

func (m *MockTransactionUseCase) Create(payload entity.Transactions) (entity.Transactions, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Transactions), args.Error(1)
}

func (m *MockTransactionUseCase) GetAll(userId string) ([]custom.TransactionsReq, error) {
	args := m.Called()
	return args.Get(0).([]custom.TransactionsReq), args.Error(1)
}

func (m *MockTransactionUseCase) GetById(id string) (custom.TransactionsReq, error) {
	args := m.Called(id)
	return args.Get(0).(custom.TransactionsReq), args.Error(1)
}
