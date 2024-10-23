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

	response := struct {
		Message string
		Data    entity.Product
	}{
		Message: "Product Created",
		Data:    Product,
	}
	c.JSON(http.StatusCreated, response)
}

func (p *ProductController) getAllProduct(c *gin.Context) {
	Products, err := p.useCase.FindAllProduct()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve data Products"})
		return
	}

	if len(Products) > 0 {
		response := struct {
			Message string
			Data    []entity.Product
		}{
			Message: "List All Product",
			Data:    Products,
		}

		c.JSON(http.StatusOK, response)
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

	response := struct {
		Message string
		Data    entity.Product
	}{
		Message: "Product found",
		Data:    Product,
	}

	c.JSON(http.StatusOK, response)
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
	response := struct {
		Message string
		Data    entity.Product
	}{
		Message: "The product has been updated",
		Data:    product,
	}

	c.JSON(http.StatusOK, response)
}

func (p *ProductController) deleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := p.useCase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	response := struct {
		Message string
		Data    entity.Product
	}{
		Message: "The product has been deleted",
		Data:    entity.Product{},
	}

	c.JSON(http.StatusNoContent, response)
}

func NewProductController(useCase usecase.ProductUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *ProductController {
	return &ProductController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
