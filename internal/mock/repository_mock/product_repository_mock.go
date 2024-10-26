package repositorymock

import (
	"server-pulsa-app/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(product entity.Product) (entity.Product, error) {
	args := m.Called(product)
	return args.Get(0).(entity.Product), args.Error(1)
}

func (m *MockProductRepository) List() ([]entity.Product, error) {
	args := m.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}

func (m *MockProductRepository) Get(id string) (entity.Product, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Product), args.Error(1)
}

func (m *MockProductRepository) Update(product entity.Product) (entity.Product, error) {
	args := m.Called(product)
	return args.Get(0).(entity.Product), args.Error(1)
}

func (m *MockProductRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
