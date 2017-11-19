package main

import (
	"os"
	"github.com/txzdream/serviceCourse/cloudgo-io/service"
)

func main() {
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "3000"
	}

	server := service.GetServer()
	server.Run(":" + PORT)
}