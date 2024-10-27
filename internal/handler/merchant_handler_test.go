package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/mock/middleware_mock"
	"server-pulsa-app/internal/mock/usecase_mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type MerchantHandlerTest struct {
	suite.Suite
	merchantUc      *usecase_mock.MerchantUsecaseMock
	router          *gin.Engine
	authMiddleware  *middleware_mock.AuthMiddlewareMock
	merchantHandler *MerchantHandler
	log             logger.Logger
}

func (m *MerchantHandlerTest) SetupTest() {
	m.merchantUc = new(usecase_mock.MerchantUsecaseMock)
	m.authMiddleware = new(middleware_mock.AuthMiddlewareMock)

	m.router = gin.Default()
	gin.SetMode(gin.TestMode)

	rg := m.router.Group("/api/v1")

	m.log = logger.NewLogger()
	m.merchantHandler = NewMerchantHandler(m.merchantUc, m.authMiddleware, rg, &m.log)
	m.router.POST("/api/v1/merchant", m.merchantHandler.createHandler)
	m.router.GET("/api/v1/merchants", m.merchantHandler.listHandler)
	m.router.GET("/api/v1/merchant/:id", m.merchantHandler.getHandler)
	m.router.PUT("/api/v1/merchant/:id", m.merchantHandler.updateHandler)
	m.router.DELETE("/api/v1/merchant/:id", m.merchantHandler.deleteHandler)
}

func (m *MerchantHandlerTest) TestCreate() {
	payload := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "Merchant Test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		m.T().Fatalf("error '%s' occured when marshaling the payload", err)
	}
	m.merchantUc.On("RegisterNewMerchant", payload).Return(payload, nil)
	request, err := http.NewRequest("POST", "/api/v1/merchant", bytes.NewBuffer(jsonPayload))
	if err != nil {
		m.T().Fatalf("error '%s' occured when creating the request", err)
	}

	w := httptest.NewRecorder()
	m.router.ServeHTTP(w, request)

	m.Equal(http.StatusCreated, w.Code)
}

func (m *MerchantHandlerTest) TestUpdate() {
	payload := entity.Merchant{
		IdMerchant:   "uuid-merchant-test",
		IdUser:       "uuid-user-test",
		NameMerchant: "Merchant Test",
		Address:      "address-test",
		IdProduct:    "uuid-product-test",
		Balance:      10000,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		m.T().Fatalf("error '%s' occured when marshaling the payload", err)
	}
	m.merchantUc.On("UpdateMerchant", payload).Return(payload, nil)
	request, err := http.NewRequest("PUT", "/api/v1/merchant/"+payload.IdMerchant, bytes.NewBuffer(jsonPayload))
	if err != nil {
		m.T().Fatalf("error '%s' occured when creating the request", err)
	}

	w := httptest.NewRecorder()
	m.router.ServeHTTP(w, request)

	m.Equal(http.StatusOK, w.Code)
}

func (m *MerchantHandlerTest) TestList() {
	m.merchantUc.On("FindAllMerchant").Return([]entity.Merchant{}, nil)
	request, err := http.NewRequest("GET", "/api/v1/merchants", nil)
	if err != nil {
		m.T().Fatalf("error '%s' occured when creating the request", err)
	}

	w := httptest.NewRecorder()
	m.router.ServeHTTP(w, request)

	m.Equal(http.StatusOK, w.Code)
}

func (m *MerchantHandlerTest) TestGet() {
	id := "uuid-merchant-test"
	m.merchantUc.On("FindMerchantByID", id).Return(entity.Merchant{}, nil)
	request, err := http.NewRequest("GET", "/api/v1/merchant/"+id, nil)
	if err != nil {
		m.T().Fatalf("error '%s' occured when creating the request", err)
	}

	w := httptest.NewRecorder()
	m.router.ServeHTTP(w, request)

	m.Equal(http.StatusOK, w.Code)
}

func (m *MerchantHandlerTest) TestDelete() {
	id := "uuid-merchant-test"
	m.merchantUc.On("DeleteMerchant", id).Return(nil)
	request, err := http.NewRequest("DELETE", "/api/v1/merchant/"+id, nil)
	if err != nil {
		m.T().Fatalf("error '%s' occured when creating the request", err)
	}

	w := httptest.NewRecorder()
	m.router.ServeHTTP(w, request)

	m.Equal(http.StatusOK, w.Code)
}

func TestMerchantHandlerSuite(t *testing.T) {
	suite.Run(t, new(MerchantHandlerTest))
}
