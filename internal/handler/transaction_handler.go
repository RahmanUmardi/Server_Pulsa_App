package handler

import (
	"net/http"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/shared/custom"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	usecase        usecase.TransactionUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

func NewTransactionHandler(usecase usecase.TransactionUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup, log *logger.Logger) *TransactionHandler {
	return &TransactionHandler{usecase: usecase, authMiddleware: authMiddleware, rg: rg, log: log}
}

func (h *TransactionHandler) createHandler(ctx *gin.Context) {
	var payload entity.Transactions

	h.log.Info("Starting to create a new transaction in the handler layer", nil)
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		h.log.Error("invalid payload for transaction", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction, err := h.usecase.Create(payload)
	if err != nil {
		h.log.Error("failed to create a transaction", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create a transaction " + err.Error()})
		return
	}
	response := struct {
		Message string              `json:"message"`
		Data    entity.Transactions `json:"data"`
	}{
		Message: "Transaction Created",
		Data:    transaction,
	}

	h.log.Info("Transaction created successfuly", response)
	ctx.JSON(http.StatusCreated, response)
}

func (h *TransactionHandler) listHandler(ctx *gin.Context) {
	h.log.Info("Starting to get transactions list in the handler layer", nil)

	transactions, err := h.usecase.GetAll()
	if err != nil {
		h.log.Error("failed to retrieve a transactions", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve transactions " + err.Error()})
		return
	}

	if len(transactions) > 0 {
		response := struct {
			Message string                   `json:"message"`
			Data    []custom.TransactionsReq `json:"data"`
		}{
			Message: "Transaction list",
			Data:    transactions,
		}
		h.log.Info("transactions list found", response)
		ctx.JSON(http.StatusOK, response)
	} else {
		h.log.Error("transactions not found", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transactions is empty"})
	}
}

func (h *TransactionHandler) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	h.log.Info("Starting to get transaction by id in the handler layer", nil)
	transaction, err := h.usecase.GetById(id)
	if err != nil {
		h.log.Error("failed to retrieve a transaction", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve a transaction" + err.Error()})
		return
	}
	response := struct {
		Message string                 `json:"message"`
		Data    custom.TransactionsReq `json:"data"`
	}{
		Message: "Transaction detail",
		Data:    transaction,
	}
	h.log.Info("transaction found", response)
	ctx.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) Route() {
	h.rg.POST(config.PostTransaction, h.authMiddleware.RequireToken("employee"), h.createHandler)
	h.rg.GET(config.ListTransactions, h.authMiddleware.RequireToken("employee"), h.listHandler)
	h.rg.GET(config.DetailTransaction, h.authMiddleware.RequireToken("employee"), h.getByIdHandler)
}
