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

var expectedMerchant = entity.Merchant{
	IdMerchant:   "uuid-merchant-test",
	IdUser:       "uuid-user-test",
	NameMerchant: "name-merchant-test",
	Address:      "address-test",
	IdProduct:    "uuid-product-test",
	Balance:      10000,
}

type merchantRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	mr      MerchantRepository
	log     logger.Logger
}

func TestMerchantRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(merchantRepositoryTestSuite))
}

func (m *merchantRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	m.NoError(err)

	m.mockDb = mockDb
	m.mockSql = mockSql
	m.log = logger.NewLogger()
	m.mr = NewMerchantRepository(mockDb, &m.log)
}

func (m *merchantRepositoryTestSuite) TestGet_success() {

	merchantRows := sqlmock.NewRows([]string{"id_merchant", "id_user", "name_merchant", "address", "id_product", "balance"}).AddRow(
		expectedMerchant.IdMerchant,
		expectedMerchant.IdUser,
		expectedMerchant.NameMerchant,
		expectedMerchant.Address,
		expectedMerchant.IdProduct,
		expectedMerchant.Balance,
	)

	m.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant WHERE id_merchant = $1")).
		WithArgs(expectedMerchant.IdMerchant).WillReturnRows(
		merchantRows,
	)

	merchant, err := m.mr.Get("uuid-merchant-test")

	m.Nil(err)
	m.Equal(expectedMerchant, merchant)
}

func (m *merchantRepositoryTestSuite) TestGet_fail() {
	m.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant WHERE id_merchant = $1")).
		WithArgs(expectedMerchant.IdMerchant).WillReturnError(sql.ErrNoRows)

	_, err := m.mr.Get("uuid-merchant-test")

	m.NotNil(err)
}

func (m *merchantRepositoryTestSuite) TestList_success() {
	merchantRows := sqlmock.NewRows([]string{"id_merchant", "id_user", "name_merchant", "address", "id_product", "balance"}).AddRow(
		expectedMerchant.IdMerchant,
		expectedMerchant.IdUser,
		expectedMerchant.NameMerchant,
		expectedMerchant.Address,
		expectedMerchant.IdProduct,
		expectedMerchant.Balance,
	)

	m.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant")).WillReturnRows(
		merchantRows,
	)

	merchants, err := m.mr.List()

	m.Nil(err)
	m.Equal([]entity.Merchant{expectedMerchant}, merchants)
}

func (m *merchantRepositoryTestSuite) TestList_fail() {
	m.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id_merchant, id_user, name_merchant, address, id_product, balance FROM mst_merchant")).WillReturnError(sql.ErrNoRows)

	_, err := m.mr.List()

	m.NotNil(err)
}

func (m *merchantRepositoryTestSuite) TestCreate_success() {
	m.mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO mst_merchant (id_user, name_merchant, address, id_product, balance) VALUES ($1, $2, $3, $4, $5) RETURNING id_merchant")).WillReturnRows(
		sqlmock.NewRows([]string{"id_merchant"}).AddRow(expectedMerchant.IdMerchant),
	)

	_, err := m.mr.Create(expectedMerchant)

	m.Nil(err)
}

func (m *merchantRepositoryTestSuite) TestCreate_fail() {
	m.mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO mst_merchant (id_merchant, id_user, name_merchant, address, id_product, balance) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_merchant")).WillReturnError(sql.ErrNoRows)

	_, err := m.mr.Create(expectedMerchant)

	m.NotNil(err)
}

func (m *merchantRepositoryTestSuite) TestDelete_fail() {
	m.mockSql.ExpectQuery(regexp.QuoteMeta("DELETE FROM mst_merchant WHERE id_merchant = $1")).WillReturnError(sql.ErrNoRows)

	err := m.mr.Delete(expectedMerchant.IdMerchant)

	m.NotNil(err)
}

func (m *merchantRepositoryTestSuite) TestUpdate_fail() {
	merchant := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test-update",
		NameMerchant: "name-merchant-test-update",
		Address:      "address-test-update",
		IdProduct:    "uuid-product-test-update",
		Balance:      20000,
	}

	m.mockSql.ExpectQuery(regexp.QuoteMeta("UPDATE mst_merchant SET id_user = $1, name_merchant = $2, address = $3, id_product = $4, balance = $5 WHERE id_merchant = $6")).WillReturnError(sql.ErrNoRows)

	_, err := m.mr.Update(merchant, expectedMerchant)

	m.NotNil(err)
}
