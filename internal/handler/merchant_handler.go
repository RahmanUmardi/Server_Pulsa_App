package handler

import (
	"net/http"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	merchantUc     usecase.MerchantUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (m *MerchantHandler) createHandler(ctx *gin.Context) {
	var payload entity.Merchant
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Invalid Payload for Merchant",
			Data:    entity.Merchant{},
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	merchant, err := m.merchantUc.RegisterNewMerchant(payload)
	if err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Merchant Creation Failed",
			Data:    entity.Merchant{},
		}
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
	ctx.JSON(http.StatusCreated, response)
}

func (m *MerchantHandler) listHandler(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, response)
}
func (m *MerchantHandler) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	merchant, err := m.merchantUc.FindMerchantByID(id)
	if err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Merchant of Id " + id + " Not Found",
			Data:    entity.Merchant{},
		}
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
	ctx.JSON(http.StatusOK, response)
}

func (m *MerchantHandler) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var payload entity.Merchant
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Invalid Payload for Merchant",
			Data:    entity.Merchant{},
		}
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
	ctx.JSON(http.StatusOK, response)
}

func (m *MerchantHandler) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := m.merchantUc.DeleteMerchant(id)
	if err != nil {
		response := struct {
			Message string
			Data    entity.Merchant
		}{
			Message: "Merchant of Id " + id + " Not Found",
			Data:    entity.Merchant{},
		}
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
	ctx.JSON(http.StatusOK, response)
}

func (m *MerchantHandler) Route() {
	m.rg.POST(config.PostMerchant, m.authMiddleware.RequireToken("employee"), m.createHandler)
	m.rg.GET(config.GetMerchantList, m.authMiddleware.RequireToken("employee"), m.listHandler)
	m.rg.GET(config.GetMerchant, m.authMiddleware.RequireToken("employee"), m.getHandler)
	m.rg.PUT(config.PutMerchant, m.authMiddleware.RequireToken("employee"), m.updateHandler)
	m.rg.DELETE(config.DeleteMerchant, m.authMiddleware.RequireToken("employee"), m.deleteHandler)
}

func NewMerchantHandler(merchantUc usecase.MerchantUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup) *MerchantHandler {
	return &MerchantHandler{merchantUc: merchantUc, authMiddleware: authMiddleware, rg: rg}
}
