package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	repositorymock "server-pulsa-app/internal/mock/repository_mock"
	"testing"

	"github.com/stretchr/testify/suite"
)

type productUsecaseTestSuite struct {
	suite.Suite
	mockProductRepository *repositorymock.MockProductRepository
	ProductUseCase        ProductUseCase
	log                   logger.Logger
}

func (p *productUsecaseTestSuite) SetupTest() {
	p.mockProductRepository = new(repositorymock.MockProductRepository)
	p.log = logger.NewLogger()
	p.ProductUseCase = NewProductUseCase(p.mockProductRepository, &p.log)
}

func (p *productUsecaseTestSuite) TestCreateNewProduct_Success() {
	newProduct := entity.Product{
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}

	CreatedProduct := entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}

	p.mockProductRepository.On("Create", newProduct).Return(CreatedProduct, nil).Once()

	product, err := p.ProductUseCase.CreateNewProduct(newProduct)

	p.Nil(err)
	p.Equal(CreatedProduct, product)
}

func (p *productUsecaseTestSuite) TestListAllProducts_Success() {
	products := []entity.Product{
		{
			IdProduct:    "1",
			NameProvider: "Product A",
			Nominal:      1000,
			Price:        1000,
			IdSupliyer:   "1",
		},
		{
			IdProduct:    "2",
			NameProvider: "Product B",
			Nominal:      2000,
			Price:        2000,
			IdSupliyer:   "2",
		},
	}

	p.mockProductRepository.On("List").Return(products, nil).Once()

	productsList, err := p.ProductUseCase.FindAllProduct()

	p.Nil(err)
	p.Equal(products, productsList)
}

func (p *productUsecaseTestSuite) TestFindProductById_Success() {
	id := "1"

	product := entity.Product{
		IdProduct:    "1",
		NameProvider: "Product A",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}

	p.mockProductRepository.On("Get", id).Return(product, nil).Once()

	productFound, err := p.ProductUseCase.FindProductById(id)

	p.Nil(err)
	p.Equal(product, productFound)
}

func (p *productUsecaseTestSuite) TestUpdateProduct_Success() {
	id := "1"

	updatedProduct := entity.Product{
		IdProduct:    "1",
		NameProvider: "Updated Product",
		Nominal:      2000,
		Price:        2000,
		IdSupliyer:   "1",
	}

	p.mockProductRepository.On("Get", id).Return(updatedProduct, nil).Once()
	p.mockProductRepository.On("Update", updatedProduct).Return(updatedProduct, nil).Once()

	productUpdated, err := p.ProductUseCase.UpdateProduct(updatedProduct)

	p.Nil(err)
	p.Equal(updatedProduct, productUpdated)
}

func (p *productUsecaseTestSuite) TestDeleteProduct_Success() {
	id := "1"

	p.mockProductRepository.On("Get", id).Return(entity.Product{}, nil).Once()
	p.mockProductRepository.On("Delete", id).Return(nil).Once()

	err := p.ProductUseCase.DeleteProduct(id)

	p.Nil(err)
}

func TestProductUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(productUsecaseTestSuite))
}
