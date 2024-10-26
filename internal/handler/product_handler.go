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

type ProductController struct {
	useCase        usecase.ProductUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

func (p *ProductController) Route() {
	p.rg.POST(config.PostProduct, p.authMiddleware.RequireToken("employee"), p.CreateProduct)
	p.rg.GET(config.GetProductList, p.authMiddleware.RequireToken("employee"), p.GetAllProduct)
	p.rg.GET(config.GetProduct, p.authMiddleware.RequireToken("employee"), p.GetProductById)
	p.rg.PUT(config.PutProduct, p.authMiddleware.RequireToken("employee"), p.UpdateProduct)
	p.rg.DELETE(config.DeleteProduct, p.authMiddleware.RequireToken("employee"), p.DeleteProduct)
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var payload entity.Product

	p.log.Info("Starting to create a new product in the handler layer", nil)

	if err := c.ShouldBindJSON(&payload); err != nil {
		p.log.Error("Invalid payload for product: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	Product, err := p.useCase.CreateNewProduct(payload)
	if err != nil {
		p.log.Error("Product creation failed", err)
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

	p.log.Info("Product created successfully", response)
	c.JSON(http.StatusCreated, response)
}

func (p *ProductController) GetAllProduct(c *gin.Context) {
	p.log.Info("Starting to retrieve all product in the handler layer", nil)

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

		p.log.Info("Product found successfully", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	p.log.Info("Product not found", nil)
	c.JSON(http.StatusOK, gin.H{"message": "List Product empty"})
}

func (p *ProductController) GetProductById(c *gin.Context) {
	id := (c.Param("id"))

	p.log.Info("Starting to retrieve product with id in the handler layer", nil)
	Product, err := p.useCase.FindProductById(id)
	if err != nil {
		p.log.Error("Product ID %s not found: ", id)
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

	p.log.Info("Product found successfully", nil)
	c.JSON(http.StatusOK, response)
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	var payload entity.Product
	id := (c.Param("id"))

	p.log.Info("Starting to update product with id in the handler layer", nil)

	if err := c.ShouldBindJSON(&payload); err != nil {
		p.log.Error("Invalid payload for product: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	payload.IdProduct = id

	p.log.Info("Updating product ID %s", id)
	product, err := p.useCase.UpdateProduct(payload)
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

	p.log.Info("Product updated successfully", response)
	c.JSON(http.StatusOK, response)
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	p.log.Info("Starting to delete product with id in the handler layer", nil)
	err := p.useCase.DeleteProduct(id)
	if err != nil {
		p.log.Error("Product ID %s not found: ", id)
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

	p.log.Info("Product deleted successfully", response)
	c.JSON(http.StatusNoContent, response)
}

func NewProductController(useCase usecase.ProductUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, log *logger.Logger) *ProductController {
	return &ProductController{useCase: useCase, rg: rg, authMiddleware: authMiddleware, log: log}
}
