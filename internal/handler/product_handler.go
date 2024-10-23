package handler

import (
	"net/http"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	useCase        usecase.ProductUseCase    // use case untuk operasi puku
	rg             *gin.RouterGroup          // router group untuk menghandle request
	authMiddleware middleware.AuthMiddleware // middleware untuk autentikasi
}

func (p *ProductController) Route() {

	p.rg.POST("/Products", p.authMiddleware.RequireToken("admin"), p.createProduct)

	p.rg.GET("/Products", p.authMiddleware.RequireToken("admin", "user"), p.getAllProduct)

	p.rg.GET("/Products/:id", p.authMiddleware.RequireToken("admin", "user"), p.getProductpyId)

	p.rg.PUT("/Products/:id", p.authMiddleware.RequireToken("admin"), p.updateProduct)

	p.rg.DELETE("/Products/:id", p.authMiddleware.RequireToken("admin"), p.deleteProduct)
}

func (p *ProductController) createProduct(c *gin.Context) {
	var payload entity.Product
	if err := c.ShouldBindJSON(&payload); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	Product, err := p.useCase.CreateNewProduct(payload)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to create Product"})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func (p *ProductController) getAllProduct(c *gin.Context) {
	Products, err := p.useCase.FindAllProduct()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve data Products"})
		return
	}

	if len(Products) > 0 {

		c.JSON(http.StatusOK, Products)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "List Product empty"})
}

func (p *ProductController) getProductpyId(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	Product, err := p.useCase.FindProductpyId(id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to get Product py ID"})
		return
	}

	c.JSON(http.StatusOK, Product)
}

// func (p *ProductController) updateProduct(c *gin.Context) {
// 	var payload entity.Product
// 	if err := c.ShouldBindJSON(&payload); err != nil {

// 		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 		return
// 	}

// 	Product, err := p.useCase.UpdateProduct(payload)
// 	if err != nil {

// 		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, Product)
// }

func (b *ProductController) updateProduct(c *gin.Context) {
	var payload entity.Product 
	id, err := uuid.Parse(c.Param("id")) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid product ID"})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	payload.IdProduct = id 

	product, err := b.useCase.UpdateProduct(payload) 
	if err != nil {
		
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product) 
}

func (p *ProductController) deleteProduct(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	err := p.useCase.DeleteProduct(id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to delete Product"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewProductController(useCase usecase.ProductUseCase, rg *gin.RouterGroup, am middleware.AuthMiddleware) *ProductController {
	return &ProductController{useCase: useCase, rg: rg, authMiddleware: am}
}
