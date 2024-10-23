package entity

type Merchant struct {
	IdMerchant   string  `json:"idMerchant"`
	IdUser       string  `json:"idUser"`
	NameMerchant string  `json:"nameMerchant"`
	Address      string  `json:"address"`
	IdProduct    string  `json:"idProduct"`
	Balance      float64 `json:"balance"`
}
