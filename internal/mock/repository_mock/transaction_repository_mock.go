package repositorymock

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/shared/custom"

	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(payload entity.Transactions) (entity.Transactions, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Transactions), args.Error(1)
}

func (m *MockTransactionRepository) GetAll(userId string) ([]custom.TransactionsReq, error) {
	args := m.Called()
	return args.Get(0).([]custom.TransactionsReq), args.Error(1)
}

func (m *MockTransactionRepository) GetById(id string) (custom.TransactionsReq, error) {
	args := m.Called(id)
	return args.Get(0).(custom.TransactionsReq), args.Error(1)
}
