package handler

import (
	"net/http"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

var logProduct = logger.GetLogger()

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

	logProduct.Info("Starting to create a new product in the handler layer")

	if err := c.ShouldBindJSON(&payload); err != nil {
		logProduct.Errorf("Invalid payload for product: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	Product, err := p.useCase.CreateNewProduct(payload)
	if err != nil {
		logProduct.Errorf("Product creation failed")
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

	logProduct.Info("Product created successfully")
	c.JSON(http.StatusCreated, response)
}

func (p *ProductController) getAllProduct(c *gin.Context) {
	logProduct.Info("Starting to retrieve all product in the handler layer")

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

		logProduct.Info("Product found successfully")
		c.JSON(http.StatusOK, response)
		return
	}

	logProduct.Info("Product not found")
	c.JSON(http.StatusOK, gin.H{"message": "List Product empty"})
}

func (p *ProductController) getProductpyId(c *gin.Context) {
	id := (c.Param("id"))

	logProduct.Info("Starting to retrieve product with id in the handler layer")
	Product, err := p.useCase.FindProductpyId(id)
	if err != nil {
		logProduct.Errorf("Product ID %s not found: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"err": "Product not found"})
		return
	}

	response := struct {
		Message string
		Data    entity.Product
	}{
		Message: "Product found",
		Data:    Product,
	}

	logProduct.Info("Product found successfully")
	c.JSON(http.StatusOK, response)
}

func (b *ProductController) updateProduct(c *gin.Context) {
	var payload entity.Product
	id := (c.Param("id"))

	logProduct.Info("Starting to update product with id in the handler layer")

	if err := c.ShouldBindJSON(&payload); err != nil {
		logProduct.Errorf("Invalid payload for product: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	payload.IdProduct = id

	logProduct.Infof("Updating product ID %s", id)
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

	logProduct.Info("Product updated successfully")
	c.JSON(http.StatusOK, response)
}

func (p *ProductController) deleteProduct(c *gin.Context) {
	id := c.Param("id")

	logProduct.Info("Starting to delete product with id in the handler layer")
	err := p.useCase.DeleteProduct(id)
	if err != nil {
		logProduct.Errorf("Product ID %s not found: %v", id, err)
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

	logProduct.Info("Product deleted successfully")
	c.JSON(http.StatusNoContent, response)
}

func NewProductController(useCase usecase.ProductUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *ProductController {
	return &ProductController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
