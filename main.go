package main

import (
	"server-pulsa-app/internal"
	"server-pulsa-app/internal/logger"
)

func main() {

	internal.NewServer().Run()
}
