package main

import (
	"fmt"
	"server-pulsa-app/internal"
)

func main() {
	fmt.Println("Server Pulsa App")
	internal.NewServer().Run()
}
