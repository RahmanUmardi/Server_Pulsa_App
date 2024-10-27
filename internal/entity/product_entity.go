package entity

type (
	Product struct {
		IdProduct    string  `db:"id_product" json:"idProduct"`
		NameProvider string  `db:"name_provider" json:"nameProvider"`
		Nominal      float64 `db:"nominal" json:"nominal"`
		Price        float64 `db:"price" json:"price"`
		IdSupliyer   string  `db:"id_supliyer" json:"idSupliyer"`
	}

	ProductRequest struct {
		NameProvider string  `json:"nameProvider" binding:"required" example:"Indosat"`
		Nominal      float64 `json:"nominal" binding:"required" example:"5000"`
		Price        float64 `json:"price" binding:"required" example:"6000"`
		IdSupliyer   string  `json:"idSupliyer" binding:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
	}

	ProductResponse struct {
		IdProduct    string  `json:"idProduct" example:"eyJhbGciOiJIUzI1NiIs..."`
		NameProvider string  `son:"nameProvider" example:"Indosat"`
		Nominal      float64 `json:"nominal" example:"5000"`
		Price        float64 `json:"price" example:"6000"`
		IdSupliyer   string  `json:"idSupliyer" example:"eyJhbGciOiJIUzI1NiIs..."`
	}

	ProductErrorResponse struct {
		Error string `json:"error" example:"Invalid product"`
	}
)
