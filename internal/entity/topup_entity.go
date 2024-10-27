package entity

import "time"

type TopupRequest struct {
	Id            string    `json:"id"`
	IdMerchant    string    `json:"id_merchant"`
	IdSupliyer    string    `json:"id_supliyer"`
	Item_name     string    `json:"item_name"`
	Amount        int       `json:"amount"`
	PaymentMethod string    `json:"va_numbers,omitempty"`
	Status        string    `json:"status,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type MidtransRequest struct {
	TransactionDetails TransactionDetails `json:"transaction_details"`
}

type TransactionDetails struct {
	OrderId     string  `json:"order_id"`
	GrossAmount float64 `json:"gross_amount"`
}

type MidtransResponse struct {
	Token       string `json:"token,omitempty"`
	RedirectURL string `json:"redirect_url"`
}

type CallbackPayment struct {
	VANumber          []VANumber `json:"va_numbers"`
	TransactionTime   string     `json:"transaction_time"`
	TransactionStatus string     `json:"transaction_status"`
	TransactionID     string     `json:"transaction_id"`
	StatusMessage     string     `json:"status_message"`
	StatusCode        string     `json:"status_code"`
	SignatureKey      string     `json:"signature_key"`
	PaymentType       string     `json:"payment_type"`
	PaymentAmount     []string   `json:"payment_amount,omitempty"`
	OrderID           string     `json:"order_id"`
	MerchantID        string     `json:"merchant_id"`
	GrossAmount       string     `json:"gross_amount"`
	FraudStatus       string     `json:"fraud_status"`
	Currency          string     `json:"currency"`
}

type VANumber struct {
	VANumber string `json:"va_number"`
	Bank     string `json:"bank"`
}

type TopupRequestDetail struct {
	Id            string    `json:"id"`
	IdMerchant    string    `json:"id_merchant"`
	IdSupliyer    Supliyer  `json:"id_supliyer"`
	Item_name     string    `json:"item_name"`
	Amount        int       `json:"amount"`
	PaymentMethod string    `json:"va_numbers,omitempty"`
	Status        string    `json:"status,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type Supliyer struct {
	IdSupliyer   string `json:"id_supliyer"`
	NameSupliyer string `json:"name_supliyer"`
}
