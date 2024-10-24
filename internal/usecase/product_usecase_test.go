package usecase

import (
	"fmt"
	"server-pulsa-app/internal/entity"
	mock "server-pulsa-app/mock/usecase_mock.go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type productUCSuite struct {
	suite.Suite
	ProductUseCase *mock.ProductUseCaseMock
	PrductUC       ProductUseCase
}

func (p *productUCSuite) SetupTest() {
	p.ProductUseCase = new(mock.ProductUseCaseMock)
}

func (p *productUCSuite) TestCreateNewProduct_Failed() {
	newProduct := entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}

	p.ProductUseCase.On("CreateNewProduct", newProduct).Return(entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}, nil).Once()

	_, err := p.PrductUC.CreateNewProduct(newProduct)

	p.ProductUseCase.AssertExpectations(p.T())

	assert.Nil(p.T(), err)

}

func (p *productUCSuite) TestCreateNewProduct_Success() {
	newProduct := entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}

	p.ProductUseCase.On("CreateNewProduct", newProduct).Return(entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}, nil).Once()

	_, err := p.PrductUC.CreateNewProduct(newProduct)

	p.ProductUseCase.AssertExpectations(p.T())

	assert.Nil(p.T(), err)

}

func (p *productUCSuite) TestFindAllProduct_Failed() {

	products := []entity.Product{
		{
			IdProduct:    "1",
			NameProvider: "Test Product",
			Nominal:      1000,
			Price:        1000,
			IdSupliyer:   "1",
		},
		{
			IdProduct:    "2",
			NameProvider: "Test Product",
			Nominal:      1000,
			Price:        1000,
			IdSupliyer:   "1",
		},
	}

	p.ProductUseCase.On("FindAllProduct").Return(products, fmt.Errorf("failed")).Once()

	_, err := p.PrductUC.FindAllProduct()

	p.ProductUseCase.AssertExpectations(p.T())

	assert.NotNil(p.T(), err)

}

func (p *productUCSuite) TestFindAllProduct_Success() {

	products := []entity.Product{
		{
			IdProduct:    "1",
			NameProvider: "Test Product",
			Nominal:      1000,
			Price:        1000,
			IdSupliyer:   "1",
		},
		{
			IdProduct:    "2",
			NameProvider: "Test Product",
			Nominal:      1000,
			Price:        1000,
			IdSupliyer:   "1",
		},
	}

	p.ProductUseCase.On("FindAllProduct").Return(products, nil).Once()

	_, err := p.PrductUC.FindAllProduct()

	p.ProductUseCase.AssertExpectations(p.T())

	assert.Nil(p.T(), err)

}

func (p *productUCSuite) TestFindProductById_Success() {

	p.ProductUseCase.On("FindProductbyId", "1").Return(entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}, nil).Once()

	_, err := p.PrductUC.FindProductById("1")

	p.ProductUseCase.AssertExpectations(p.T())

	assert.Nil(p.T(), err)

}

func (p *productUCSuite) TestFindProductById_Failed() {

	p.ProductUseCase.On("FindProductbyId", "1").Return(entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}, nil).Once()

	_, err := p.PrductUC.FindProductById("1")

	p.ProductUseCase.AssertExpectations(p.T())

	assert.NotNil(p.T(), err)

}
func (p *productUCSuite) TestUpdateProduct_Success() {

	newProduct := entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}

	p.ProductUseCase.On("UpdateProduct", newProduct).Return(entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}, nil).Once()

	_, err := p.PrductUC.UpdateProduct(newProduct)

	p.ProductUseCase.AssertExpectations(p.T())

	assert.Nil(p.T(), err)

}

func (p *productUCSuite) TestUpdateProduct_Failed() {

	newProduct := entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}

	p.ProductUseCase.On("UpdateProduct", newProduct).Return(entity.Product{
		IdProduct:    "1",
		NameProvider: "Test Product",
		Nominal:      1000,
		Price:        1000,
		IdSupliyer:   "1",
	}, fmt.Errorf("failed")).Once()

	_, err := p.PrductUC.UpdateProduct(newProduct)

	p.ProductUseCase.AssertExpectations(p.T())

	assert.NotNil(p.T(), err)

}

func (p *productUCSuite) TestDeleteProduct_Success() {

	p.ProductUseCase.On("DeleteProduct", "1").Return(nil).Once()

	err := p.PrductUC.DeleteProduct("1")

	p.ProductUseCase.AssertExpectations(p.T())

	assert.Nil(p.T(), err)

}

func (p *productUCSuite) TestDeleteProduct_Failed() {

	p.ProductUseCase.On("DeleteProduct", "1").Return(fmt.Errorf("failed")).Once()

	err := p.PrductUC.DeleteProduct("1")

	p.ProductUseCase.AssertExpectations(p.T())

	assert.NotNil(p.T(), err)

}
