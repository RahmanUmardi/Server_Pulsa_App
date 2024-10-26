package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	am "server-pulsa-app/internal/mock/auth_mock"
	mock "server-pulsa-app/internal/mock/usecase_mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TransactionHandlerTestSuite struct {
	suite.Suite
	mockTxUc           *mock.MockTransactionUseCase
	mockAuthMiddleware *am.AuthMiddlewareMock
	TransactionHandler *TransactionHandler
	router             *gin.Engine
	log                logger.Logger
}

func (suite *TransactionHandlerTestSuite) SetupTest() {
	suite.mockTxUc = new(mock.MockTransactionUseCase)
	suite.mockAuthMiddleware = new(am.AuthMiddlewareMock)
	gin.SetMode(gin.TestMode)
	suite.router = gin.New()

	suite.TransactionHandler = NewTransactionHandler(suite.mockTxUc, suite.mockAuthMiddleware, suite.router.Group("/api/v1"), &suite.log)
	suite.router.POST("/transaction", suite.TransactionHandler.createHandler)
	suite.router.GET("/transactions/history", suite.TransactionHandler.listHandler)
	suite.router.GET("/transactions/history/:id", suite.TransactionHandler.getByIdHandler)
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
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	suite.mockTxUc.On("Create", payload).Return(payload, nil)

	req, err := http.NewRequest("POST", "/api/v1/transaction", bytes.NewBuffer(jsonPayload))

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

}

func (suite *TransactionHandlerTestSuite) TestGetAll_Success() {

	suite.mockTxUc.On("GetAll").Return([]entity.Transactions{}, nil)

	req, err := http.NewRequest("GET", "/api/v1/transactions/history", nil)

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}

func (suite *TransactionHandlerTestSuite) TestGetById_Success() {
	id := "uuid-test"

	suite.mockTxUc.On("GetById", id).Return(entity.Transactions{}, nil)

	req, err := http.NewRequest("GET", "/api/v1/transaction/history/"+id, nil)

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}

func TestTransactionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionHandlerTestSuite))
}
