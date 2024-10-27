package main

import (
	_ "server-pulsa-app/docs"
	"server-pulsa-app/internal"
)

// @title Server Pulsa API
// @version 1.0
// @description API Server for Pulsa Application

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
func main() {

	internal.NewServer().Run()
}
