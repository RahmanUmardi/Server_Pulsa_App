package entity

type (
	Transactions struct {
		TransactionsId    string              `json:"transactionId"`
		MerchantId        string              `json:"merchantId"`
		UserId            string              `json:"userId"`
		CustomerName      string              `json:"customerName"`
		DestinationNumber string              `json:"destinationNumber"`
		TransactionDate   string              `json:"transactionDate"`
		TransactionDetail []TransactionDetail `json:"transactionDetail"`
	}

	TransactionDetail struct {
		TransactionDetailId string  `json:"transactionDetailId"`
		TransactionsId      string  `json:"transactionId"`
		ProductId           string  `json:"productId"`
		Price               float64 `json:"Price"`
	}

	TransactionReq struct {
		MerchantId        string                 `json:"merchantId" binding:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
		UserId            string                 `json:"userId" binding:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
		CustomerName      string                 `json:"customerName" binding:"required" example:"customer a"`
		DestinationNumber string                 `json:"destinationNumber" binding:"required" example:"08...."`
		TransactionDate   string                 `json:"transactionDate" binding:"required" example:"27-10-2024"`
		TransactionDetail []TransactionDetailReq `json:"transactionDetail" binding:"required"`
	}

	TransactionDetailReq struct {
		ProductId string `json:"productId" binding:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
	}

	TransactionErrorResponse struct {
		Error string `json:"error" example:"Invalid transaction"`
	}
)
