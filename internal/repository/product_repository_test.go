package repository

import (
	"database/sql"
	"regexp"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type productRepoTestSuite struct {
	suite.Suite
	mockDB      *sql.DB
	mockSql     sqlmock.Sqlmock
	productRepo ProductRepository
	log         logger.Logger
}

func (p *productRepoTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		p.T().Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}

	p.mockDB = mockDb
	p.mockSql = mockSql
	p.log = logger.NewLogger()
	p.productRepo = NewProductRepository(p.mockDB, &p.log)
}

func (p *productRepoTestSuite) TearDownTest() {
	p.mockDB.Close()
}

func (p *productRepoTestSuite) TestCreateProduct_Repository() {
	product := entity.Product{
		NameProvider: "Provider A",
		Nominal:      10000,
		Price:        12000,
		IdSupliyer:   "Supplier A",
	}

	query := "INSERT INTO mst_product (name_provider, nominal, price, id_supliyer) VALUES ($1, $2, $3, $4) RETURNING id_product"

	p.mockSql.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(product.NameProvider, product.Nominal, product.Price, product.IdSupliyer).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdProduct, err := p.productRepo.Create(product)

	p.Nil(err)
	p.Equal("1", createdProduct.IdProduct)
	p.Equal(product.NameProvider, createdProduct.NameProvider)
	p.Equal(product.Nominal, createdProduct.Nominal)
	p.Equal(product.Price, createdProduct.Price)
	p.Equal(product.IdSupliyer, createdProduct.IdSupliyer)
}

func (p *productRepoTestSuite) TestGetProductById_Repository() {
	id := "1"

	query := "SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product WHERE id_product = $1"

	p.mockSql.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id_product", "name_provider", "nominal", "price", "id_supliyer"}).AddRow(id, "Provider A", 10000, 12000, "Supplier A"))

	product, err := p.productRepo.Get(id)

	p.Nil(err)
	p.Equal("1", product.IdProduct)
	p.Equal("Provider A", product.NameProvider)
	p.Equal(float64(10000), product.Nominal)
	p.Equal(float64(12000), product.Price)
	p.Equal("Supplier A", product.IdSupliyer)
}

func (p *productRepoTestSuite) TestFindAllProduct_Repository() {
	query := "SELECT id_product, name_provider, nominal, price, id_supliyer FROM mst_product"

	p.mockSql.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(sqlmock.NewRows([]string{"id_product", "name_provider", "nominal", "price", "id_supliyer"}).
		AddRow("1", "Provider A", 10000, 12000, "Supplier A").
		AddRow("2", "Provider B", 20000, 24000, "Supplier B"))

	products, err := p.productRepo.List()

	p.Nil(err)
	p.Len(products, 2)
	p.Equal("1", products[0].IdProduct)
	p.Equal("Provider A", products[0].NameProvider)
	p.Equal(float64(10000), products[0].Nominal)
	p.Equal(float64(12000), products[0].Price)
	p.Equal("Supplier A", products[0].IdSupliyer)
	p.Equal("2", products[1].IdProduct)
	p.Equal("Provider B", products[1].NameProvider)
	p.Equal(float64(20000), products[1].Nominal)
	p.Equal(float64(24000), products[1].Price)
	p.Equal("Supplier B", products[1].IdSupliyer)
}

func (p *productRepoTestSuite) TestUpdateProduct_Repository() {
	product := entity.Product{
		IdProduct:    "1",
		NameProvider: "Provider A",
		Nominal:      10000,
		Price:        12000,
		IdSupliyer:   "Supplier A",
	}

	query := "UPDATE mst_product SET name_provider = $1, nominal = $2, price = $3, id_supliyer = $4 WHERE id_product = $5"

	p.mockSql.ExpectExec(regexp.QuoteMeta(query)).WithArgs(product.NameProvider, product.Nominal, product.Price, product.IdSupliyer, product.IdProduct).WillReturnResult(sqlmock.NewResult(1, 1))

	updatedProduct, err := p.productRepo.Update(product)

	p.Nil(err)
	p.Equal("1", updatedProduct.IdProduct)
	p.Equal("Provider A", updatedProduct.NameProvider)
	p.Equal(float64(10000), updatedProduct.Nominal)
	p.Equal(float64(12000), updatedProduct.Price)
	p.Equal("Supplier A", updatedProduct.IdSupliyer)
}

func (p *productRepoTestSuite) TestDeleteProduct_Repository() {
	id := "1"

	query := "DELETE FROM mst_product WHERE id_product = $1"

	p.mockSql.ExpectExec(regexp.QuoteMeta(query)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := p.productRepo.Delete(id)

	p.Nil(err)
}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(productRepoTestSuite))
}
