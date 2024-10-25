package main

import (
	"server-pulsa-app/internal"
	"server-pulsa-app/internal/logger"
)

func main() {
	logger.InitLogger()
	log := logger.GetLogger()
	log.Info("Server Pulsa App Started")
	internal.NewServer().Run()
}
