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

	//transaction route
	PostTransaction   = "/transaction"
	ListTransactions  = "/transactions"
	DetailTransaction = "/transaction/:id"

	// user route
	GetUserList = "/users"
	GetUser     = "/user/:id"
	PutUser     = "/user/:id"
	DeleteUser  = "/user/:id"

	// auth route
	Login    = "/auth/login"
	Register = "/auth/register"

	// topup route
	PostTopup            = "/topup"
	GetTopupByMerchantId = "/topup/:id"
	// GetTopupList = "/topups"

	// callback topup
	PostCallback = "/topup/callback"

	//report route
	GetReport = "/report"
)
