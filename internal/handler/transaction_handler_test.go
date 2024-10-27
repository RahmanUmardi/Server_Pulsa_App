package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	am "server-pulsa-app/internal/mock/auth_mock"
	mock "server-pulsa-app/internal/mock/usecase_mock"
	"server-pulsa-app/internal/shared/custom"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TransactionHandlerTestSuite struct {
	suite.Suite
	mockTxUc           *mock.MockTransactionUseCase
	mockAuthMiddleware *am.AuthMiddlewareMock
	transactionHandler *TransactionHandler
	router             *gin.Engine
	log                logger.Logger
}

func (suite *TransactionHandlerTestSuite) SetupTest() {
	suite.mockTxUc = new(mock.MockTransactionUseCase)
	suite.mockAuthMiddleware = new(am.AuthMiddlewareMock)
	gin.SetMode(gin.TestMode)
	suite.router = gin.New()

	rg := suite.router.Group("/api/v1")
	suite.transactionHandler = NewTransactionHandler(suite.mockTxUc, suite.mockAuthMiddleware, rg, &suite.log)
	suite.transactionHandler.Route()
}

func (suite *TransactionHandlerTestSuite) TestCreate_Success() {
	payload := entity.Transactions{
		MerchantId:        "uuid-test1",
		UserId:            "uuid-test1",
		CustomerName:      "test",
		DestinationNumber: "087654321",
		TransactionDate:   "25-10-2024",
		TransactionDetail: []entity.TransactionDetail{
			{
				ProductId: "uuid-test",
				Price:     50000,
			},
		},
	}

	expectedResponse := entity.Transactions{
		TransactionsId:    "tx-uuid",
		MerchantId:        "uuid-test1",
		UserId:            "uuid-test1",
		CustomerName:      "test",
		DestinationNumber: "087654321",
		TransactionDate:   "25-10-2024",
		TransactionDetail: []entity.TransactionDetail{
			{
				TransactionDetailId: "detail-uuid",
				TransactionsId:      "tx-uuid",
				ProductId:           "uuid-test",
				Price:               50000,
			},
		},
	}

	suite.mockTxUc.On("Create", payload).Return(expectedResponse, nil)

	jsonPayload, err := json.Marshal(payload)
	suite.NoError(err)

	req, err := http.NewRequest("POST", "/api/v1/transaction", bytes.NewBuffer(jsonPayload))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

	var response struct {
		Message string              `json:"message"`
		Data    entity.Transactions `json:"data"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("Transaction Created", response.Message)
	suite.Equal(expectedResponse, response.Data)
}

func (suite *TransactionHandlerTestSuite) TestCreate_InvalidPayload() {
	invalidPayload := `{"invalid": "json`

	req, err := http.NewRequest("POST", "/api/v1/transaction", bytes.NewBufferString(invalidPayload))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestCreate_UseCaseError() {
	payload := entity.Transactions{
		MerchantId:        "uuid-test1",
		UserId:            "uuid-test1",
		CustomerName:      "test",
		DestinationNumber: "087654321",
		TransactionDate:   "25-10-2024",
		TransactionDetail: []entity.TransactionDetail{
			{
				ProductId: "uuid-test",
				Price:     50000,
			},
		},
	}

	suite.mockTxUc.On("Create", payload).Return(entity.Transactions{}, errors.New("usecase error"))

	jsonPayload, err := json.Marshal(payload)
	suite.NoError(err)

	req, err := http.NewRequest("POST", "/api/v1/transaction", bytes.NewBuffer(jsonPayload))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestGetAll_Success() {
	expectedTransactions := []custom.TransactionsReq{
		{
			TransactionsId:    "tx-uuid",
			CustomerName:      "test",
			DestinationNumber: "087654321",
			TransactionDate:   time.Now(),
			User: custom.UserRes{
				Id_user:  "user-uuid",
				Username: "testuser",
				Role:     "employee",
			},
			Merchant: custom.MerchantRes{
				IdMerchant:   "merchant-uuid",
				NameMerchant: "Test Merchant",
				Address:      "Test Address",
			},
			TransactionDetail: []custom.TransactionDetailReq{
				{
					TransactionDetailId: "detail-uuid",
					TransactionsId:      "tx-uuid",
					Product: custom.ProductRes{
						IdProduct:    "product-uuid",
						NameProvider: "Test Provider",
						Nominal:      50000,
						Price:        55000,
					},
				},
			},
		},
	}

	suite.mockTxUc.On("GetAll").Return(expectedTransactions, nil)

	req, err := http.NewRequest("GET", "/api/v1/transactions/history", nil)
	suite.NoError(err)

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response struct {
		Message string                   `json:"message"`
		Data    []custom.TransactionsReq `json:"data"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("Transaction list", response.Message)
	suite.Equal(expectedTransactions, response.Data)
}

func (suite *TransactionHandlerTestSuite) TestGetAll_Empty() {
	suite.mockTxUc.On("GetAll").Return([]custom.TransactionsReq{}, nil)

	req, err := http.NewRequest("GET", "/api/v1/transactions/history", nil)
	suite.NoError(err)

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestGetAll_Error() {
	suite.mockTxUc.On("GetAll").Return([]custom.TransactionsReq{}, errors.New("usecase error"))

	req, err := http.NewRequest("GET", "/api/v1/transactions/history", nil)
	suite.NoError(err)

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *TransactionHandlerTestSuite) TestGetById_Success() {
	id := "tx-uuid"
	expectedTransaction := custom.TransactionsReq{
		TransactionsId:    id,
		CustomerName:      "test",
		DestinationNumber: "087654321",
		TransactionDate:   time.Now(),
		User: custom.UserRes{
			Id_user:  "user-uuid",
			Username: "testuser",
			Role:     "employee",
		},
		Merchant: custom.MerchantRes{
			IdMerchant:   "merchant-uuid",
			NameMerchant: "Test Merchant",
			Address:      "Test Address",
		},
		TransactionDetail: []custom.TransactionDetailReq{
			{
				TransactionDetailId: "detail-uuid",
				TransactionsId:      id,
				Product: custom.ProductRes{
					IdProduct:    "product-uuid",
					NameProvider: "Test Provider",
					Nominal:      50000,
					Price:        55000,
				},
			},
		},
	}

	suite.mockTxUc.On("GetById", id).Return(expectedTransaction, nil)

	req, err := http.NewRequest("GET", "/api/v1/transactions/history/"+id, nil)
	suite.NoError(err)

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response struct {
		Message string                 `json:"message"`
		Data    custom.TransactionsReq `json:"data"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("Transaction detail", response.Message)
	suite.Equal(expectedTransaction, response.Data)
}

func (suite *TransactionHandlerTestSuite) TestGetById_Error() {
	id := "non-existent-id"
	suite.mockTxUc.On("GetById", id).Return(custom.TransactionsReq{}, errors.New("usecase error"))

	req, err := http.NewRequest("GET", "/api/v1/transactions/history/"+id, nil)
	suite.NoError(err)

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusInternalServerError, w.Code)
}

func TestTransactionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionHandlerTestSuite))
}
