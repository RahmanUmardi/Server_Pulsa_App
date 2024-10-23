package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/repository"

	"github.com/google/uuid"
)

type ProductUseCase interface {
	CreateNewProduct(Product entity.Product) (entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
	FindProductpyId(id uuid.UUID) (entity.Product, error)
	UpdateProduct(Product entity.Product) (entity.Product, error)
	DeleteProduct(id uuid.UUID) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(Product entity.Product) (entity.Product, error) {
	return p.repo.Create(Product)
}

func (p *productUseCase) FindAllProduct() ([]entity.Product, error) {
	return p.repo.List()
}

func (p *productUseCase) FindProductpyId(id uuid.UUID) (entity.Product, error) {
	return p.repo.Get(id)
}

func (p *productUseCase) UpdateProduct(product entity.Product) (entity.Product, error) {
	_, err := p.repo.Get(product.IdProduct)
	if err != nil {
		return entity.Product{}, fmt.Errorf("Product with ID %d not found", product.IdProduct)
	}
	return p.repo.Update(product.IdProduct,product)
}

func (p *productUseCase) DeleteProduct(id uuid.UUID) error {
	_, err := p.repo.Get(id)
	if err != nil {
		return fmt.Errorf("Product with ID %d not found", id)
	}
	return p.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}
