package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/repository"
)

// var logProduct = logger.GetLogger()

type ProductUseCase interface {
	CreateNewProduct(Product entity.Product) (entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
	FindProductById(id string) (entity.Product, error)
	UpdateProduct(Product entity.Product) (entity.Product, error)
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo repository.ProductRepository
	log  *logger.Logger
}

func (p *productUseCase) CreateNewProduct(Product entity.Product) (entity.Product, error) {
	p.log.Info("Starting to create a new product in the usecase layer", nil)
	return p.repo.Create(Product)
}

func (p *productUseCase) FindAllProduct() ([]entity.Product, error) {
	p.log.Info("Starting to retrive all product in the usecase layer", nil)
	return p.repo.List()
}

func (p *productUseCase) FindProductById(id string) (entity.Product, error) {
	p.log.Info("Starting to retrive a product by id in the usecase layer", nil)
	return p.repo.Get(id)
}

func (p *productUseCase) UpdateProduct(product entity.Product) (entity.Product, error) {
	p.log.Info("Starting to retrive a product by id in the usecase layer", nil)

	_, err := p.repo.Get(product.IdProduct)
	if err != nil {
		return entity.Product{}, fmt.Errorf("product with ID %s not found", product.IdProduct)
	}

	p.log.Info("Product ID %s has been updated successfully: ", product.IdProduct)
	return p.repo.Update(product)
}

func (p *productUseCase) DeleteProduct(id string) error {
	p.log.Info("Starting to retrive a product by id in the usecase layer", nil)

	_, err := p.repo.Get(id)
	if err != nil {
		return fmt.Errorf("product with ID %s not found", id)
	}

	p.log.Info("Product has been deleted successfully: ", id)
	return p.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository, log *logger.Logger) ProductUseCase {
	return &productUseCase{repo: repo, log: log}
}
