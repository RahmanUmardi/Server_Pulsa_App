package main

import (
	"server-pulsa-app/internal"
	"server-pulsa-app/internal/logger"
)

func main() {
	logger.InitLogger()
	log := logger.GetLogger()
	log.Info("Server Pulsa App")
	internal.NewServer().Run()
}
