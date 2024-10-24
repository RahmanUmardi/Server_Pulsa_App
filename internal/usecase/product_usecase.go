package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/repository"
)

var logProduct = logger.GetLogger()

type ProductUseCase interface {
	CreateNewProduct(Product entity.Product) (entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
	FindProductById(id string) (entity.Product, error)
	UpdateProduct(Product entity.Product) (entity.Product, error)
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) CreateNewProduct(Product entity.Product) (entity.Product, error) {
	logProduct.Info("Starting to create a new product in the usecase layer")
	return p.repo.Create(Product)
}

func (p *productUseCase) FindAllProduct() ([]entity.Product, error) {
	logProduct.Info("Starting to retrive all product in the usecase layer")
	return p.repo.List()
}

func (p *productUseCase) FindProductpyId(id string) (entity.Product, error) {
	logProduct.Info("Starting to retrive a product by id in the usecase layer")
	return p.repo.Get(id)
}

func (p *productUseCase) UpdateProduct(product entity.Product) (entity.Product, error) {
	logProduct.Info("Starting to retrive a product by id in the usecase layer")

	_, err := p.repo.Get(product.IdProduct)
	if err != nil {
		logProduct.Errorf("Product ID %s not found: %v", product.IdProduct, err)
		return entity.Product{}, fmt.Errorf("product with ID %s not found", product.IdProduct)
	}

	logProduct.Infof("Product ID %s has been updated successfully: ", product.IdProduct)
	return p.repo.Update(product.IdProduct, product)
}

func (p *productUseCase) DeleteProduct(id string) error {
	logProduct.Info("Starting to retrive a product by id in the usecase layer")

	_, err := p.repo.Get(id)
	if err != nil {
		logProduct.Errorf("Product ID %s not found: %v", id, err)
		return fmt.Errorf("product with ID %s not found", id)
	}

	logProduct.Info("Product has been deleted successfully: ", id)
	return p.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}
