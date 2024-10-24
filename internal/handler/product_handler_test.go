package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockProductUseCase struct{
	
}

func (m *mockProductUseCase) CreateNewProduct(product entity.Product) (entity.Product, error) {
	return product, nil
}

func (m *mockProductUseCase) FindAllProduct() ([]entity.Product, error) {
	return []entity.Product{}, nil
}

func (m *mockProductUseCase) FindProductById(id string) (entity.Product, error) {
	return entity.Product{IdProduct: id}, nil
}

func (m *mockProductUseCase) UpdateProduct(product entity.Product) (entity.Product, error) {
	return product, nil
}

func (m *mockProductUseCase) DeleteProduct(id string) error {
	return nil
}

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	useCase := &mockProductUseCase{}
	authMiddleware := middleware.NewAuthMiddleware()
	productController := NewProductController(useCase, router.Group("/products"), authMiddleware)

	router.POST("/products", productController.createProduct)

	product := entity.Product{IdProduct: "1", NameProvider: "Test Product"}
	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer token")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestGetAllProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	useCase := &mockProductUseCase{}
	authMiddleware := middleware.AuthMiddleware{}
	productController := handler.NewProductController(useCase, router.Group("/products"), authMiddleware)

	router.GET("/products", productController.GetAllProduct)

	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set("Authorization", "Bearer token")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetProductById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	useCase := &mockProductUseCase{}
	authMiddleware := middleware.AuthMiddleware{}
	productController := handler.NewProductController(useCase, router.Group("/products"), authMiddleware)

	router.GET("/products/:id", productController.GetProductById)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	req.Header.Set("Authorization", "Bearer token")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUpdateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	useCase := &mockProductUseCase{}
	authMiddleware := middleware.AuthMiddleware{}
	productController := handler.NewProductController(useCase, router.Group("/products"), authMiddleware)

	router.PUT("/products/:id", productController.UpdateProduct)

	product := entity.Product{IdProduct: "1", NameProvider: "Updated Product"}
	jsonValue, _ := json.Marshal(product)

	req, _ := http.NewRequest("PUT", "/products/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer token")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	useCase := &mockProductUseCase{}
	authMiddleware := middleware.AuthMiddleware{}
	productController := handler.NewProductController(useCase, router.Group("/products"), authMiddleware)

	router.DELETE("/products/:id", productController.DeleteProduct)

	req, _ := http.NewRequest("DELETE", "/products/1", nil)
	req.Header.Set("Authorization", "Bearer token")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNoContent, res.Code)
}

// package handler

// import (
// 	"bytes"
// 	"server-pulsa-app/internal/entity"
// 	"server-pulsa-app/internal/middleware"
// 	mock "server-pulsa-app/mock/usecase_mock.go"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gin-gonic/gin/internal/json"
// 	"github.com/stretchr/testify/suite"
// )

// type ProductControllerTest struct {
// 	suite.Suite
// 	rg        *gin.Engine
// 	productUC *mock.ProductUseCaseMock
// 	// am middleware.AuthMiddleware
// }

// // func (s *ProductControllerTest) SetupTest() {
// // 	s.productUC = new(mock.ProductUseCaseMock)
// // 	s.rg = gin.Default()

// // 	gin.SetMode(gin.TestMode)

// // 	rg := s.rg.Group("/api/v1")

// // 	productUC := NewProductController(s.productUC,rg,am)
// // 	s.rg.po
// // }

// func (s *ProductControllerTest) TestCreateProductHandler_success() {
// 	payload := entity.Product{
// 		IdProduct:    "1",
// 		NameProvider: "test",
// 		Nominal:      10000,
// 		Price:        10000,
// 		IdSupliyer:   "1",
// 	}

// 	s.productUC.On("CreateNewProduct", payload).Return(entity.Product{
// 		IdProduct:    "1",
// 		NameProvider: "test",
// 		Nominal:      10000,
// 		Price:        10000,
// 		IdSupliyer:   "1",
// 	}, nil).Once()

// 	var buf bytes.Buffer
// 	err:= json.NewEncoder(&buf).NewEncoder(payload)

// 	if err != nil {
// 		panic(err)
// 	}

// }

// // func (s *ProductControllerTest) TestCreateProductHandler_success() {
// // 	payload := entity.Product{
// // 		IdProduct:    "1",
// // 		NameProvider: "test",
// // 		Nominal:      10000,
// // 		Price:        10000,
// // 		IdSupliyer:   "1",
// // 	}

// // 	s.productUC.On("CreateNewProduct", payload).Return(entity.Product{
// // 		IdProduct:    "1",
// // 		NameProvider: "test",
// // 		Nominal:      10000,
// // 		Price:        10000,
// // 		IdSupliyer:   "1",
// // 	}, nil).Once()

// // 	var buf bytes.Buffer
// // 	err:= json.NewEncoder(&buf).NewEncoder(payload)

// // 	if err != nil {
// // 		panic(err)
// // 	}
