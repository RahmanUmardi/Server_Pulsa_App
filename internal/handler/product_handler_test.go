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

type ProductControllerTestSuite struct {
	suite.Suite
	mockProductUC      *mock.ProductUseCaseMock
	mockAuthMiddleware *am.AuthMiddlewareMock
	ProductController  *ProductController
	router             *gin.Engine
	log                logger.Logger
}

func (suite *ProductControllerTestSuite) SetupTest() {
	suite.mockProductUC = new(mock.ProductUseCaseMock)
	suite.mockAuthMiddleware = new(am.AuthMiddlewareMock)
	gin.SetMode(gin.TestMode)
	suite.router = gin.New()

	suite.log = logger.NewLogger()
	suite.ProductController = NewProductController(suite.mockProductUC, suite.router.Group("/api/v1/products"), suite.mockAuthMiddleware, &suite.log)
	suite.router.POST("/api/v1/product", suite.ProductController.CreateProduct)
	suite.router.PUT("/api/v1/product/:id", suite.ProductController.UpdateProduct)
	suite.router.DELETE("/api/v1/product/:id", suite.ProductController.DeleteProduct)
	suite.router.GET("/api/v1/products", suite.ProductController.GetAllProduct)
	suite.router.GET("/api/v1/product/:id", suite.ProductController.GetProductById)
}

func (suite *ProductControllerTestSuite) TestCreateProduct() {
	payload := entity.Product{
		IdProduct:    "1",
		NameProvider: "Axis",
		Nominal:      10000,
		Price:        11000,
		IdSupliyer:   "1",
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	suite.mockProductUC.On("CreateNewProduct", payload).Return(payload, nil)

	req, err := http.NewRequest("POST", "/api/v1/product", bytes.NewBuffer(jsonPayload))

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

}

func (suite *ProductControllerTestSuite) TestGetProductById() {
	id := "1"
	intID := "1"

	suite.mockProductUC.On("FindProductById", intID).Return(entity.Product{}, nil)

	req, err := http.NewRequest("GET", "/api/v1/product/"+id, nil)

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}

func (suite *ProductControllerTestSuite) TestUpdateProduct() {
	payload := entity.Product{
		IdProduct:    "1",
		NameProvider: "Axis",
		Nominal:      10000,
		Price:        11000,
		IdSupliyer:   "1",
	}

	suite.mockProductUC.On("UpdateProduct", payload).Return(payload, nil)

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", "/api/v1/product/1", bytes.NewBuffer(jsonPayload))

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}

func (suite *ProductControllerTestSuite) TestDeleteProduct() {
	id := "1"
	intID := "1"

	suite.mockProductUC.On("DeleteProduct", intID).Return(nil)

	req, err := http.NewRequest("DELETE", "/api/v1/product/"+id, nil)

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNoContent, w.Code)

}

func (suite *ProductControllerTestSuite) TestGetAllProduct() {

	suite.mockProductUC.On("FindAllProduct").Return([]entity.Product{}, nil)

	req, err := http.NewRequest("GET", "/api/v1/products", nil)

	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}

func TestProductControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductControllerTestSuite))
}
