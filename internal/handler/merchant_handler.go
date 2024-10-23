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
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	merchant, err := m.merchantUc.RegisterNewMerchant(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, merchant)
}

func (m *MerchantHandler) listHandler(ctx *gin.Context) {
	merchants, err := m.merchantUc.FindAllMerchant()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(merchants) > 0 {
		ctx.JSON(http.StatusOK, merchants)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "List of merchant is empty"})
}
func (m *MerchantHandler) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	merchant, err := m.merchantUc.FindMerchantByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, merchant)
}

func (m *MerchantHandler) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var payload entity.Merchant
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	payload.IdMerchant = id

	expense, err := m.merchantUc.UpdateMerchant(payload)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, expense)
}

func (m *MerchantHandler) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := m.merchantUc.DeleteMerchant(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
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
