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
)
