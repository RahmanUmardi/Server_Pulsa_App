package handler

import (
	"net/http"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	useCase        usecase.ProductUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (p *ProductController) Route() {
	p.rg.POST(config.PostProduct, p.authMiddleware.RequireToken("employee"), p.createProduct)
	p.rg.GET(config.GetProductList, p.authMiddleware.RequireToken("employee"), p.getAllProduct)
	p.rg.GET(config.GetProduct, p.authMiddleware.RequireToken("employee"), p.getProductpyId)
	p.rg.PUT(config.PutProduct, p.authMiddleware.RequireToken("employee"), p.updateProduct)
	p.rg.DELETE(config.DeleteProduct, p.authMiddleware.RequireToken("employee"), p.deleteProduct)
}

func (p *ProductController) createProduct(c *gin.Context) {
	var payload entity.Product
	if err := c.ShouldBindJSON(&payload); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	Product, err := p.useCase.CreateNewProduct(payload)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
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
	id := (c.Param("id"))
	Product, err := p.useCase.FindProductpyId(id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to get Product py ID"})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func (b *ProductController) updateProduct(c *gin.Context) {
	var payload entity.Product
	id := (c.Param("id"))

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
	id := c.Param("id")

	err := p.useCase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewProductController(useCase usecase.ProductUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *ProductController {
	return &ProductController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
