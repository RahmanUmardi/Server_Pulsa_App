package handler

import (
	"server-pulsa-app/config"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportUc       usecase.ReportUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

// ListMerchantReport godoc
// @Summary transaction report
// @Description Download the transaction report
// @Tags transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} custom.ReportRes "Transaction report downloaded"
// @Failure 401 {object} entity.TransactionErrorResponse "Unauthorized"
// @Router /report [get]
func (r *ReportHandler) listHandler(ctx *gin.Context) {
	r.log.Info("Starting to retrieve all merchant's transactions in the handler layer", nil)

	userId, _ := ctx.Get("employee")
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	r.reportUc.FindAllMerchant(userId.(string), startDate, endDate)
	// r.reportUc.FindAllMerchant(userId.(string))
	// if err != nil {
	// 	response := struct {
	// 		Message string
	// 		Data    []custom.TransactionsReq
	// 	}{
	// 		Message: err.Error(),
	// 		Data:    nil,
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, response)
	// 	return
	// }

	// if len(report) > 0 {
	// 	response := struct {
	// 		Message string
	// 		Data    []custom.TransactionsReq
	// 	}{
	// 		Message: "Merchant List Found",
	// 		Data:    report,
	// 	}

	// 	r.log.Info("Merchant found successfully", nil)
	// 	ctx.JSON(http.StatusOK, response)
	// 	return
	// }
	// response := struct {
	// 	Message string
	// 	Data    []custom.TransactionsReq
	// }{
	// 	Message: "List of merchant is empty",
	// 	Data:    nil,
	// }

	// r.log.Info("Merchant not found", response)
	// ctx.JSON(http.StatusOK, response)
}

func (m *ReportHandler) Route() {
	m.rg.GET(config.GetReport, m.authMiddleware.RequireToken("employee"), m.listHandler)
}

func NewReportHandler(reportUc usecase.ReportUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup, log *logger.Logger) *ReportHandler {
	return &ReportHandler{reportUc: reportUc, authMiddleware: authMiddleware, rg: rg, log: log}
}
