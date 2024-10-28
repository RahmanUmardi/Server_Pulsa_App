package handler

import (
	"net/http"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/shared/custom"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportUc       usecase.ReportUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

func (r *ReportHandler) listHandler(ctx *gin.Context) {
	r.log.Info("Starting to retrieve all merchant's transactions in the handler layer", nil)

	userId, _ := ctx.Get("employee")
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	err := r.reportUc.FindAllTransactions(userId.(string), startDate, endDate)
	if err != nil {
		response := struct {
			Message string
			Data    []custom.TransactionsReq
		}{
			Message: err.Error(),
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := struct {
		Message string
	}{
		Message: "Report as Excel File Generated Successfully",
	}
	r.log.Info("Report as Excel File Generated Successfully", nil)
	ctx.JSON(http.StatusOK, response)
}

func (m *ReportHandler) Route() {
	m.rg.GET(config.GetReport, m.authMiddleware.RequireToken("employee"), m.listHandler)
}

func NewReportHandler(reportUc usecase.ReportUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup, log *logger.Logger) *ReportHandler {
	return &ReportHandler{reportUc: reportUc, authMiddleware: authMiddleware, rg: rg, log: log}
}
