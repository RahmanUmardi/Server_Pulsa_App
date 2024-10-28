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

// @title Merchant API
// @version 1.0
// @description Merchant management endpoints for the server-pulsa-app
type MerchantHandler struct {
	merchantUc     usecase.MerchantUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

// CreateMerchant godoc
// @Summary Create new merchant
// @Description Create a new merchant in the system
// @Tags merchants
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.MerchantRequest true "Merchant details"
// @Success 201 {object} entity.MerchantResponse "Successfully created"
// @Failure 400 {object} entity.MerchantErrorResponse "Invalid input"
// @Failure 401 {object} entity.MerchantErrorResponse "Unauthorized"
// @Router /merchant [post]
func (m *MerchantHandler) createHandler(ctx *gin.Context) {
	var payload entity.Merchant

	m.log.Info("Starting to create a new merchant in the handler layer", nil)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Invalid Payload for Merchant",
			Data:    entity.Merchant{},
		}

		m.log.Error("Invalid payload for merchant: ", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	merchant, err := m.merchantUc.RegisterNewMerchant(payload)
	if err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: err.Error(),
			Data:    entity.Merchant{},
		}

		m.log.Error("Merchant creation failed", response)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := struct {
		Message string
		Data    entity.Merchant
	}{
		Message: "Merchant Created",
		Data:    merchant,
	}

	m.log.Info("Merchant created successfully", response)
	ctx.JSON(http.StatusCreated, response)
}

// ListMerchants godoc
// @Summary List all merchants
// @Description Get a list of all merchants
// @Tags merchants
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {array} []entity.MerchantResponse "List of merchants"
// @Failure 401 {object} entity.MerchantErrorResponse "Unauthorized"
// @Router /merchants [get]
func (m *MerchantHandler) listHandler(ctx *gin.Context) {
	m.log.Info("Starting to retrieve all merchant in the handler layer", nil)

	merchants, err := m.merchantUc.FindAllMerchant()
	if err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: err.Error(),
			Data:    entity.Merchant{},
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(merchants) > 0 {
		response := struct {
			Message string
			Data    []entity.Merchant
		}{
			Message: "Merchant List Found",
			Data:    merchants,
		}

		m.log.Info("Merchant found successfully", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := struct {
		Message string
		Data    entity.Merchant
	}{
		Message: "List of merchant is empty",
		Data:    entity.Merchant{},
	}

	m.log.Info("Merchant not found", response)
	ctx.JSON(http.StatusOK, response)
}

// GetMerchant godoc
// @Summary Get merchant by ID
// @Description Retrieve a merchant by its ID
// @Tags merchants
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Merchant ID"
// @Success 200 {object} entity.MerchantResponse "Merchant found"
// @Failure 404 {object} entity.MerchantErrorResponse "Merchant not found"
// @Failure 401 {object} entity.MerchantErrorResponse "Unauthorized"
// @Router /merchant/{id} [get]
func (m *MerchantHandler) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	m.log.Info("Starting to retrieve merchant with id in the handler layer", nil)
	merchant, err := m.merchantUc.FindMerchantByID(id)
	if err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: err.Error(),
			Data:    entity.Merchant{},
		}

		m.log.Error("Merchant ID %s not found: ", response)
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := struct {
		Message string
		Data    entity.Merchant
	}{
		Message: "Merchant Found",
		Data:    merchant,
	}

	m.log.Info("Merchant found successfully", nil)
	ctx.JSON(http.StatusOK, response)
}

// UpdateMerchant godoc
// @Summary Update merchant
// @Description Update an existing merchant
// @Tags merchants
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Merchant ID"
// @Param request body entity.MerchantRequest true "Updated merchant details"
// @Success 200 {object} entity.MerchantResponse "Successfully updated merchant"
// @Failure 400 {object} entity.MerchantErrorResponse "Invalid input"
// @Failure 401 {object} entity.MerchantErrorResponse "Unauthorized"
// @Failure 404 {object} entity.MerchantErrorResponse "Merchant not found"
// @Router /merchant/{id} [put]
func (m *MerchantHandler) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var payload entity.Merchant

	m.log.Info("Starting to update merchant with id in the handler layer", nil)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Invalid Payload for Merchant",
			Data:    entity.Merchant{},
		}

		m.log.Error("Invalid payload for merchant: ", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	payload.IdMerchant = id

	merchant, err := m.merchantUc.UpdateMerchant(payload)
	if err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Merchant of Id " + id + " Not Found",
			Data:    entity.Merchant{},
		}

		m.log.Error("Merchant ID %s not found: ", response)
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := struct {
		Message string
		Data    entity.Merchant
	}{
		Message: "Merchant of Id " + id + " Updated",
		Data:    merchant,
	}

	m.log.Info("Merchant updated successfully", response)
	ctx.JSON(http.StatusOK, response)
}

// DeleteMerchant godoc
// @Summary Delete merchant
// @Description Delete a merchant by its ID
// @Tags merchants
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Merchant ID"
// @Success 204 "Successfully deleted"
// @Failure 401 {object} entity.MerchantErrorResponse "Unauthorized"
// @Failure 404 {object} entity.MerchantErrorResponse "Merchant not found"
// @Router /merchant/{id} [delete]
func (m *MerchantHandler) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	m.log.Info("Starting to delete merchant with id in the handler layer", nil)
	err := m.merchantUc.DeleteMerchant(id)
	if err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Merchant of Id " + id + " Not Found",
			Data:    entity.Merchant{},
		}

		m.log.Error("Merchant ID %s not found: ", response)
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	response := struct {
		Message string
	}{
		Message: "Merchant of Id " + id + " Deleted",
	}

	m.log.Info("Merchant deleted successfully", response)
	ctx.JSON(http.StatusOK, response)
}

func (m *MerchantHandler) Route() {
	m.rg.POST(config.PostMerchant, m.authMiddleware.RequireToken("admin"), m.createHandler)
	m.rg.GET(config.GetMerchantList, m.authMiddleware.RequireToken("admin"), m.listHandler)
	m.rg.GET(config.GetMerchant, m.authMiddleware.RequireToken("admin"), m.getHandler)
	m.rg.PUT(config.PutMerchant, m.authMiddleware.RequireToken("admin"), m.updateHandler)
	m.rg.DELETE(config.DeleteMerchant, m.authMiddleware.RequireToken("admin"), m.deleteHandler)
}

func NewMerchantHandler(merchantUc usecase.MerchantUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup, log *logger.Logger) *MerchantHandler {
	return &MerchantHandler{merchantUc: merchantUc, authMiddleware: authMiddleware, rg: rg, log: log}
}
