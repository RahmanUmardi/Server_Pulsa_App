package handler

import (
	"net/http"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MerchantHandler struct {
	merchantUc     usecase.MerchantUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

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
func (m *MerchantHandler) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	m.log.Info("Starting to retrieve merchant with id in the handler layer", nil)
	merchant, err := m.merchantUc.FindMerchantByID(id)
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
		Message: "Merchant Found",
		Data:    merchant,
	}

	m.log.Info("Merchant found successfully", nil)
	ctx.JSON(http.StatusOK, response)
}

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
		Data    entity.Merchant
	}{
		Message: "Merchant of Id " + id + " Deleted",
		Data:    entity.Merchant{},
	}

	m.log.Info("Merchant deleted successfully", response)
	ctx.JSON(http.StatusOK, response)
}

func (m *MerchantHandler) Route() {
	m.rg.POST(config.PostMerchant, m.authMiddleware.RequireToken("employee"), m.createHandler)
	m.rg.GET(config.GetMerchantList, m.authMiddleware.RequireToken("employee"), m.listHandler)
	m.rg.GET(config.GetMerchant, m.authMiddleware.RequireToken("employee"), m.getHandler)
	m.rg.PUT(config.PutMerchant, m.authMiddleware.RequireToken("employee"), m.updateHandler)
	m.rg.DELETE(config.DeleteMerchant, m.authMiddleware.RequireToken("employee"), m.deleteHandler)
}

func NewMerchantHandler(merchantUc usecase.MerchantUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup, log *logger.Logger) *MerchantHandler {
	return &MerchantHandler{merchantUc: merchantUc, authMiddleware: authMiddleware, rg: rg, log: log}
}
