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

// @title Product API
// @version 1.0
// @description Product management endpoints for the server-pulsa-app

type ProductController struct {
	useCase        usecase.ProductUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

func (p *ProductController) Route() {
	p.rg.POST(config.PostProduct, p.authMiddleware.RequireToken("admin"), p.CreateProduct)
	p.rg.GET(config.GetProductList, p.authMiddleware.RequireToken("admin"), p.GetAllProduct)
	p.rg.GET(config.GetProduct, p.authMiddleware.RequireToken("admin"), p.GetProductById)
	p.rg.PUT(config.PutProduct, p.authMiddleware.RequireToken("admin"), p.UpdateProduct)
	p.rg.DELETE(config.DeleteProduct, p.authMiddleware.RequireToken("admin"), p.DeleteProduct)
}

// CreateProduct godoc
// @Summary Create new product
// @Description Create a new product in the system
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.ProductRequest true "Product details"
// @Success 201 {object} entity.ProductResponse "Successfully created product"
// @Failure 400 {object} entity.ProductErrorResponse "Invalid input"
// @Failure 401 {object} entity.ProductErrorResponse "Unauthorized"
// @Router /product [post]
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

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} []entity.ProductResponse "List of products"
// @Failure 401 {object} entity.ProductErrorResponse "Unauthorized"
// @Router /products [get]
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

// GetProduct godoc
// @Summary Get product by ID
// @Description Retrieve a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} entity.ProductResponse "Product found"
// @Failure 404 {object} entity.ProductErrorResponse "Product not found"
// @Failure 401 {object} entity.ProductErrorResponse "Unauthorized"
// @Router /product/{id} [get]
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

// UpdateProduct godoc
// @Summary Update product
// @Description Update an existing product
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Param request body entity.ProductRequest true "Updated product details"
// @Success 200 {object} entity.ProductResponse "Successfully updated product"
// @Failure 400 {object} entity.ProductErrorResponse "Invalid input"
// @Failure 401 {object} entity.ProductErrorResponse "Unauthorized"
// @Failure 404 {object} entity.ProductErrorResponse "Product not found"
// @Router /product/{id} [put]
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

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 204 "Successfully deleted"
// @Failure 401 {object} entity.ProductErrorResponse "Unauthorized"
// @Failure 404 {object} entity.ProductErrorResponse "Product not found"
// @Router /product/{id} [delete]
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
