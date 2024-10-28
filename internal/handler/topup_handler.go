package handler

import (
	"fmt"
	"os"
	"server-pulsa-app/config"
	"server-pulsa-app/internal/entity"
	"server-pulsa-app/internal/logger"
	"server-pulsa-app/internal/middleware"
	"server-pulsa-app/internal/shared/common"
	"server-pulsa-app/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type TopupHandler struct {
	usecase        usecase.TopupUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
	log            *logger.Logger
}

func initRestyClient() *resty.Client {
	client := resty.New()
	client.SetBaseURL(os.Getenv("BASE_URL_MIDTRANS"))
	client.SetHeader("Authorization", "Basic "+os.Getenv("SERVER_KEY_MIDTRANS"))
	return client
}

func (t *TopupHandler) CreateTopup(c *gin.Context) {
	var payload entity.TopupRequest

	t.log.Info("Starting to create a new topup in the handler layer", nil)
	if err := c.ShouldBindJSON(&payload); err != nil {
		t.log.Error("Invalid payload for topup: ", err)
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	payload.Status = "pending"

	t.log.Info("Start validating if the top-up amount is less than 10,000", payload.Amount)
	if payload.Amount < 10000 {
		t.log.Error("Invalid topup amount", payload.Amount)
		common.SendErrorResponse(c, 400, "minimum amount for topup is 10000")
		return
	}

	t.log.Info("Starting validate if merchant and supliyer exist", nil)
	if payload.IdMerchant == "" || payload.IdSupliyer == "" || payload.Item_name == "" {
		t.log.Error("id_merchant, id_supliyer, and item_name are required", nil)
		common.SendErrorResponse(c, 400, "id_merchant and id_supliyer are required")
		return
	}

	t.log.Info("Starting to send a payload to the usecase layer", nil)
	id, err := t.usecase.CreateTopup(payload)
	if err != nil {
		t.log.Error("Topup creation failed", err)
		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	client := initRestyClient()
	t.log.Info("Starting to send a payload to Midtrans", nil)
	midtransReq := entity.MidtransRequest{
		TransactionDetails: entity.TransactionDetails{
			OrderId:     id,
			GrossAmount: float64(payload.Amount),
		},
	}

	fmt.Printf("Payload yang dikirim ke Midtrans: %+v\n", midtransReq)

	resp, err := client.R().
		SetBody(midtransReq).
		SetResult(&entity.MidtransResponse{}).
		Post("")

	if err != nil {
		t.log.Error("Error sending payload to Midtrans: ", err)
		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	t.log.Info("Starting to validate status code", nil)
	if resp.StatusCode() != 201 {
		common.SendErrorResponse(c, resp.StatusCode(), resp.String())
		return
	}

	midtransResponse := resp.Result().(*entity.MidtransResponse)
	t.log.Info("Request topup successfully", midtransResponse)
	common.SendSingleResponseCreated(c, midtransResponse, "Please make a balance payment at the link above using the virtual account payment method from BCA, BRI, or BNI")
}

func (t *TopupHandler) PaymentCallbackHandler(c *gin.Context) {
	var notifPayment entity.CallbackPayment

	t.log.Info("Starting to handle payment callback", nil)
	if err := c.ShouldBindJSON(&notifPayment); err != nil {
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	t.log.Info("Get the data needed for the update", nil)
	idTopup := notifPayment.OrderID
	status := notifPayment.TransactionStatus
	amount, err := strconv.ParseFloat(notifPayment.GrossAmount, 64)
	if err != nil {
		common.SendErrorResponse(c, 400, err.Error())
		return
	}
	var paymentMethod string
	if len(notifPayment.VANumber) > 0 {
		paymentMethod = notifPayment.VANumber[0].Bank
	}

	t.log.Info("Update the status and payment method if status condition = settlement", nil)
	if status == "settlement" {
		payload := entity.TopupRequest{
			Id:            idTopup,
			Status:        status,
			Amount:        int(amount),
			PaymentMethod: paymentMethod,
		}

		t.log.Info("Starting to update the topup data", nil)
		idTopupSuccess, err := t.usecase.UpdateAfterPayment(payload)
		if err != nil {
			t.log.Error("Error updating topup data: ", err)
			common.SendErrorResponse(c, 500, err.Error())
			return
		}

		t.log.Info("Topup data updated successfully", nil)
		common.SendSingleResponseOk(c, idTopupSuccess, "Topup berhasil")
		return
	}

	t.log.Info("Topup status is not settlement", nil)
	common.SendErrorResponse(c, 400, "Topup gagal")

}

func (t *TopupHandler) GetTopupByMerchantId(c *gin.Context) {
	idMerchant := c.Param("id")

	if idMerchant == "" {
		t.log.Error("id_merchant is required", nil)
		common.SendErrorResponse(c, 400, "id_merchant is required")
		return
	}

	t.log.Info("Starting to get topup by merchant id", nil)
	topups, err := t.usecase.GetTopupByMerchantId(idMerchant)
	if err != nil {
		t.log.Error("Error getting topup by merchant id: ", err)
		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	t.log.Info("Topup data retrieved successfully", nil)
	common.SendSingleResponseOk(c, topups, "Data topup")
}

func (t *TopupHandler) Route() {
	t.rg.POST(config.PostTopup, t.authMiddleware.RequireToken("admin"), t.CreateTopup)
	t.rg.POST(config.PostCallback, t.PaymentCallbackHandler)
	t.rg.GET(config.GetTopupByMerchantId, t.authMiddleware.RequireToken("admin"), t.GetTopupByMerchantId)
}

func NewTopupHandler(usecase usecase.TopupUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup, log *logger.Logger) *TopupHandler {
	return &TopupHandler{usecase: usecase, authMiddleware: authMiddleware, rg: rg, log: log}
}
