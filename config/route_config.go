package config

const (
	ApiGroup = "/api/v1"
	// merchant route
	PostMerchant    = "/merchant"
	GetMerchantList = "/merchants"
	GetMerchant     = "/merchant/:id"
	PutMerchant     = "/merchant/:id"
	DeleteMerchant  = "/merchant/:id"

	// product route
	PostProduct    = "/product"
	GetProductList = "/products"
	GetProduct     = "/product/:id"
	PutProduct     = "/product/:id"
	DeleteProduct  = "/product/:id"
)
