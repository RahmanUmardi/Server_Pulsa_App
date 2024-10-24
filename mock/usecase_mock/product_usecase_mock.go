package mock

import (
	"server-pulsa-app/internal/entity"

	"github.com/stretchr/testify/mock"
)

type ProductUseCaseMock struct {
	mock.Mock
}

// Create adalah mock dari metode Create
func (m *ProductUseCaseMock) Create(product entity.Product) (entity.Product, error) {
	args := m.Called(product)
	return args.Get(0).(entity.Product), args.Error(1)
}

// List adalah mock dari metode List
func (m *ProductUseCaseMock) List() ([]entity.Product, error) {
	args := m.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}

// Get adalah mock dari metode Get
func (m *ProductUseCaseMock) Get(id string) (entity.Product, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Product), args.Error(1)
}

// Update adalah mock dari metode Update
func (m *ProductUseCaseMock) Update(id string, product entity.Product) (entity.Product, error) {
	args := m.Called(id, product)
	return args.Get(0).(entity.Product), args.Error(1)
}

// Delete adalah mock dari metode Delete
func (m *ProductUseCaseMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}