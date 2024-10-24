package mock

import (
	"github.com/stretchr/testify/mock"
	"server-pulsa-app/internal/entity"
)

// ProductRepository adalah mock dari ProductRepository
type ProductRepository struct {
	mock.Mock
}

// Create adalah mock dari metode Create
func (m *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	args := m.Called(product)
	return args.Get(0).(entity.Product), args.Error(1)
}

// List adalah mock dari metode List
func (m *ProductRepository) List() ([]entity.Product, error) {
	args := m.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}

// Get adalah mock dari metode Get
func (m *ProductRepository) Get(id string) (entity.Product, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Product), args.Error(1)
}

// Update adalah mock dari metode Update
func (m *ProductRepository) Update(id string, product entity.Product) (entity.Product, error) {
	args := m.Called(id, product)
	return args.Get(0).(entity.Product), args.Error(1)
}

// Delete adalah mock dari metode Delete
func (m *ProductRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}