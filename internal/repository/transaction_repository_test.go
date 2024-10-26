package repository

import (
	"database/sql"
	"regexp"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/shared/custom"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

var (
	expectedTransaction = entity.Transactions{
		TransactionsId:    "test-uuid",
		MerchantId:        "merchant-uuid",
		UserId:            "user-uuid",
		CustomerName:      "John",
		DestinationNumber: "081234567890",
		TransactionDate:   "25-10-2024",
		TransactionDetail: []entity.TransactionDetail{
			{
				TransactionDetailId: "detail-uuid",
				TransactionsId:      "test-uuid",
				ProductId:           "product-uuid",
				Price:               50000,
			},
		},
	}

	expectedTransactionReq = custom.TransactionsReq{
		TransactionsId:    "test-uuid",
		CustomerName:      "John Doe",
		DestinationNumber: "081234567890",
		TransactionDate:   time.Now(),
		User: custom.UserRes{
			Id_user:  "user-uuid",
			Username: "testuser",
			Role:     "admin",
		},
		Merchant: custom.MerchantRes{
			IdMerchant:   "merchant-uuid",
			NameMerchant: "Test Merchant",
			Address:      "Test Address",
		},
		TransactionDetail: []custom.TransactionDetailReq{
			{
				TransactionDetailId: "detail-uuid",
				TransactionsId:      "test-uuid",
				Product: custom.ProductRes{
					IdProduct:    "product-uuid",
					NameProvider: "Test Provider",
					Nominal:      50000,
					Price:        50000,
				},
			},
		},
	}
)

type transactionRepositoryTestSuite struct {
	suite.Suite
	mockDb          *sql.DB
	mockSql         sqlmock.Sqlmock
	transactionRepo TransactionRepository
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(transactionRepositoryTestSuite))
}

func (s *transactionRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	s.NoError(err)

	s.mockDb = mockDb
	s.mockSql = mockSql
	s.transactionRepo = NewTransactionRepository(mockDb)
}

func (s *transactionRepositoryTestSuite) TearDownTest() {
	s.mockDb.Close()
}

func (s *transactionRepositoryTestSuite) TestCreate_Success() {
	// Mock merchant existence check
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM mst_merchant WHERE id_merchant = $1)`)).
		WithArgs(expectedTransaction.MerchantId).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	// Mock user existence check
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM mst_user WHERE id_user = $1)`)).
		WithArgs(expectedTransaction.UserId).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	// Mock transaction begin
	s.mockSql.ExpectBegin()

	// Mock transaction insert
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions`)).
		WithArgs(
			expectedTransaction.MerchantId,
			expectedTransaction.UserId,
			expectedTransaction.CustomerName,
			expectedTransaction.DestinationNumber,
			sqlmock.AnyArg(), // For the parsed date
		).
		WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow(expectedTransaction.TransactionsId))

	// Mock product price query
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM mst_product WHERE id_product = $1)`)).
		WithArgs(expectedTransaction.TransactionDetail[0].ProductId).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT price FROM mst_product WHERE id_product = $1`)).
		WithArgs(expectedTransaction.TransactionDetail[0].ProductId).
		WillReturnRows(sqlmock.NewRows([]string{"price"}).AddRow(50000))

	// Mock transaction detail insert
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transaction_detail`)).
		WithArgs(
			expectedTransaction.TransactionsId,
			expectedTransaction.TransactionDetail[0].ProductId,
			50000,
		).
		WillReturnRows(sqlmock.NewRows([]string{"transaction_detail_id"}).AddRow("detail-uuid"))

	// Mock commit
	s.mockSql.ExpectCommit()

	result, err := s.transactionRepo.Create(expectedTransaction)

	s.NoError(err)
	s.Equal(expectedTransaction.TransactionsId, result.TransactionsId)
	s.Equal(expectedTransaction.CustomerName, result.CustomerName)
}

func (s *transactionRepositoryTestSuite) TestCreate_InvalidDate() {
	invalidTransaction := expectedTransaction
	invalidTransaction.TransactionDate = "invalid-date"

	result, err := s.transactionRepo.Create(invalidTransaction)

	s.Error(err)
	s.Contains(err.Error(), "invalid date format")
	s.Equal(entity.Transactions{}, result)
}

func (s *transactionRepositoryTestSuite) TestCreate_MerchantNotFound() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM mst_merchant WHERE id_merchant = $1)`)).
		WithArgs(expectedTransaction.MerchantId).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	result, err := s.transactionRepo.Create(expectedTransaction)

	s.Error(err)
	s.Equal("merchant not found", err.Error())
	s.Equal(entity.Transactions{}, result)
}

// GetAll Tests
func (s *transactionRepositoryTestSuite) TestGetAll_Success() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
		WillReturnRows(sqlmock.NewRows([]string{
			"transaction_id", "customer_name", "destination_number", "transaction_date",
			"id_user", "username", "role",
			"id_merchant", "name_merchant", "address",
			"transaction_detail_id", "transaction_id", "id_product", "name_provider", "nominal", "price",
		}).AddRow(
			expectedTransactionReq.TransactionsId,
			expectedTransactionReq.CustomerName,
			expectedTransactionReq.DestinationNumber,
			expectedTransactionReq.TransactionDate,
			expectedTransactionReq.User.Id_user,
			expectedTransactionReq.User.Username,
			expectedTransactionReq.User.Role,
			expectedTransactionReq.Merchant.IdMerchant,
			expectedTransactionReq.Merchant.NameMerchant,
			expectedTransactionReq.Merchant.Address,
			expectedTransactionReq.TransactionDetail[0].TransactionDetailId,
			expectedTransactionReq.TransactionsId,
			expectedTransactionReq.TransactionDetail[0].Product.IdProduct,
			expectedTransactionReq.TransactionDetail[0].Product.NameProvider,
			expectedTransactionReq.TransactionDetail[0].Product.Nominal,
			expectedTransactionReq.TransactionDetail[0].Product.Price,
		))

	result, err := s.transactionRepo.GetAll()

	s.NoError(err)
	s.Len(result, 1)
	s.Equal(expectedTransactionReq.TransactionsId, result[0].TransactionsId)
}

func (s *transactionRepositoryTestSuite) TestGetAll_EmptyResult() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
		WillReturnRows(sqlmock.NewRows([]string{
			"transaction_id", "customer_name", "destination_number", "transaction_date",
			"id_user", "username", "role",
			"id_merchant", "name_merchant", "address",
			"transaction_detail_id", "transaction_id", "id_product", "name_provider", "nominal", "price",
		}))

	result, err := s.transactionRepo.GetAll()

	s.NoError(err)
	s.Empty(result)
}

// GetById Tests
func (s *transactionRepositoryTestSuite) TestGetById_Success() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
		WithArgs(expectedTransactionReq.TransactionsId).
		WillReturnRows(sqlmock.NewRows([]string{
			"transaction_id", "customer_name", "destination_number", "transaction_date",
			"id_user", "username", "role",
			"id_merchant", "name_merchant", "address",
			"transaction_detail_id", "id_product", "name_provider", "nominal", "price",
		}).AddRow(
			expectedTransactionReq.TransactionsId,
			expectedTransactionReq.CustomerName,
			expectedTransactionReq.DestinationNumber,
			expectedTransactionReq.TransactionDate,
			expectedTransactionReq.User.Id_user,
			expectedTransactionReq.User.Username,
			expectedTransactionReq.User.Role,
			expectedTransactionReq.Merchant.IdMerchant,
			expectedTransactionReq.Merchant.NameMerchant,
			expectedTransactionReq.Merchant.Address,
			expectedTransactionReq.TransactionDetail[0].TransactionDetailId,
			expectedTransactionReq.TransactionDetail[0].Product.IdProduct,
			expectedTransactionReq.TransactionDetail[0].Product.NameProvider,
			expectedTransactionReq.TransactionDetail[0].Product.Nominal,
			expectedTransactionReq.TransactionDetail[0].Product.Price,
		))

	result, err := s.transactionRepo.GetById(expectedTransactionReq.TransactionsId)

	s.NoError(err)
	s.Equal(expectedTransactionReq.TransactionsId, result.TransactionsId)
}

func (s *transactionRepositoryTestSuite) TestGetById_NotFound() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
		WithArgs("non-existent-id").
		WillReturnRows(sqlmock.NewRows([]string{
			"transaction_id", "customer_name", "destination_number", "transaction_date",
			"id_user", "username", "role",
			"id_merchant", "name_merchant", "address",
			"transaction_detail_id", "id_product", "name_provider", "nominal", "price",
		}))

	result, err := s.transactionRepo.GetById("non-existent-id")

	s.Error(err)
	s.Equal("transaction not found", err.Error())
	s.Equal(custom.TransactionsReq{}, result)
}
