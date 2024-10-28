package entity

type (
	Merchant struct {
		IdMerchant   string  `json:"idMerchant"`
		IdUser       string  `json:"idUser"`
		NameMerchant string  `json:"nameMerchant"`
		Address      string  `json:"address"`
		IdProduct    string  `json:"idProduct"`
		Balance      float64 `json:"balance"`
	}

	MerchantRequest struct {
		IdUser       string `json:"idUser" binding:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
		NameMerchant string `json:"nameMerchant" binding:"required" example:"Konter Pak Eko"`
		Address      string `json:"address" binding:"required" example:"Jombang"`
		IdProduct    string `json:"idProduct" binding:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
	}

	MerchantResponse struct {
		IdMerchant   string  `json:"idMerchant" example:"eyJhbGciOiJIUzI1NiIs..."`
		IdUser       string  `json:"idUser" example:"eyJhbGciOiJIUzI1NiIs..."`
		NameMerchant string  `json:"nameMerchant" example:"Toko Pak Eko"`
		Address      string  `json:"address" example:"Jombang"`
		IdProduct    string  `json:"idProduct" example:"eyJhbGciOiJIUzI1NiIs..."`
		Balance      float64 `json:"balance" example:"500000"`
	}

	MerchantErrorResponse struct {
		Error string `json:"error" example:"Invalid merchant"`
	}
)
