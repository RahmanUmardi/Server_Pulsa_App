package custom

import "time"

type (
	TransactionsReq struct {
		TransactionsId    string                 `json:"transactionId"`
		CustomerName      string                 `json:"customerName"`
		DestinationNumber string                 `json:"destinationNumber"`
		User              UserRes                `json:"user"`
		Merchant          MerchantRes            `json:"merchant"`
		TransactionDate   time.Time              `json:"transactionDate"`
		TransactionDetail []TransactionDetailReq `json:"transactionDetail"`
	}

	TransactionDetailReq struct {
		TransactionDetailId string     `json:"transactionDetailId"`
		TransactionsId      string     `json:"transactionId,omitempty"`
		Product             ProductRes `json:"product"`
	}

	UserRes struct {
		Id_user  string `json:"id_user"`
		Username string `json:"name"`
		Role     string `json:"role"`
	}

	MerchantRes struct {
		IdMerchant   string `json:"idMerchant"`
		NameMerchant string `json:"nameMerchant"`
		Address      string `json:"address"`
	}

	ProductRes struct {
		IdProduct    string  ` json:"idProduct"`
		NameProvider string  ` json:"nameProvider"`
		Nominal      float64 ` json:"nominal"`
		Price        float64 ` json:"price"`
	}
)
