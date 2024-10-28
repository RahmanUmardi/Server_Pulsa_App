package usecase

import (
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	repositorymock "server-pulsa-app/internal/mock/repository_mock"
	"server-pulsa-app/internal/shared/custom"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type transactionUsecaseTestSuite struct {
	suite.Suite
	mockTransactionRepo *repositorymock.MockTransactionRepository
	transactionUseCase  TransactionUseCase
	log                 logger.Logger
}

func (tx *transactionUsecaseTestSuite) SetupTest() {
	tx.mockTransactionRepo = new(repositorymock.MockTransactionRepository)
	tx.log = logger.NewLogger()
	tx.transactionUseCase = NewTransactionUseCase(tx.mockTransactionRepo, &tx.log)
}

func (tx *transactionUsecaseTestSuite) TestCreate_Success() {
	newTx := entity.Transactions{
		MerchantId:        "uuid-test",
		UserId:            "uuid-test",
		CustomerName:      "custtest",
		DestinationNumber: "087654321",
		TransactionDate:   "25-10-2024",
		TransactionDetail: []entity.TransactionDetail{
			{
				ProductId: "uuid-test",
			},
		},
	}

	CreatedTx := entity.Transactions{

		TransactionsId:    "uuid-test",
		MerchantId:        "uuid-test",
		UserId:            "uuid-test",
		CustomerName:      "custtest",
		DestinationNumber: "087654321",
		TransactionDate:   "25-10-2024",
		TransactionDetail: []entity.TransactionDetail{
			{
				TransactionDetailId: "uuid-test",
				TransactionsId:      "uuid-test",
				ProductId:           "uuid-test",
				Price:               6000,
			},
		},
	}

	tx.mockTransactionRepo.On("Create", newTx).Return(CreatedTx, nil).Once()

	transaction, err := tx.transactionUseCase.Create(newTx)

	tx.Nil(err)
	tx.Equal(CreatedTx, transaction)
}

func (tx *transactionUsecaseTestSuite) TestList_Success() {
	parsedDate, err := time.Parse(time.RFC3339, "2024-10-25T00:00:00Z")
	tx.Require().NoError(err)

	transactions := []custom.TransactionsReq{
		{
			TransactionsId:    "uuid-test",
			CustomerName:      "custtest",
			DestinationNumber: "087654321",
			User: custom.UserRes{
				Id_user:  "uuid-test",
				Username: "unametest",
				Role:     "roletest",
			},
			Merchant: custom.MerchantRes{
				IdMerchant:   "uuid-test",
				NameMerchant: "nametest",
				Address:      "addresstest",
			},
			TransactionDate: parsedDate,
			TransactionDetail: []custom.TransactionDetailReq{
				{
					TransactionDetailId: "uuid-test",
					TransactionsId:      "uuid-test",
					Product: custom.ProductRes{
						IdProduct:    "uui-test",
						NameProvider: "test",
						Nominal:      5000,
						Price:        6000,
					},
				},
			},
		},
		{
			TransactionsId:    "uuid-test2",
			CustomerName:      "custtest2",
			DestinationNumber: "087654321",
			User: custom.UserRes{
				Id_user:  "uuid-test2",
				Username: "unametest2",
				Role:     "roletest2",
			},
			Merchant: custom.MerchantRes{
				IdMerchant:   "uuid-test2",
				NameMerchant: "nametest2",
				Address:      "addresstest2",
			},
			TransactionDate: parsedDate,
			TransactionDetail: []custom.TransactionDetailReq{
				{
					TransactionDetailId: "uuid-test2",
					TransactionsId:      "uuid-test2",
					Product: custom.ProductRes{
						IdProduct:    "uui-test2",
						NameProvider: "test2",
						Nominal:      5000,
						Price:        6000,
					},
				},
			},
		},
	}

	tx.mockTransactionRepo.On("List").Return(transactions, nil).Once()

	txList, err := tx.transactionUseCase.GetAll("")

	tx.Nil(err)
	tx.Equal(transactions, txList)
}

func (tx *transactionUsecaseTestSuite) TestGetById_Success() {
	id := "uuid-test 1"

	parsedDate, err := time.Parse(time.RFC3339, "2024-10-25T00:00:00Z")
	tx.Require().NoError(err)

	transaction := custom.TransactionsReq{
		TransactionsId:    "uuid-test",
		CustomerName:      "custtest",
		DestinationNumber: "087654321",
		User: custom.UserRes{
			Id_user:  "uuid-test",
			Username: "unametest",
			Role:     "roletest",
		},
		Merchant: custom.MerchantRes{
			IdMerchant:   "uuid-test",
			NameMerchant: "nametest",
			Address:      "addresstest",
		},
		TransactionDate: parsedDate,
		TransactionDetail: []custom.TransactionDetailReq{
			{
				TransactionDetailId: "uuid-test",
				TransactionsId:      "uuid-test",
				Product: custom.ProductRes{
					IdProduct:    "uui-test",
					NameProvider: "test",
					Nominal:      5000,
					Price:        6000,
				},
			},
		},
	}

	tx.mockTransactionRepo.On("Get", id).Return(transaction, nil).Once()

	txFound, err := tx.transactionUseCase.GetById(id)

	tx.Nil(err)
	tx.Equal(transaction, txFound)
}

func TestTransactionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(transactionUsecaseTestSuite))
}
