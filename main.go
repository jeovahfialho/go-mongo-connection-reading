package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	client, err := InitDB() // Importa a função InitDB de database.go
	if err != nil {
		fmt.Println("Error initializing the database:", err)
		return
	}
	defer client.Disconnect(context.TODO())

	SetupRoutes(router, client) // Importa a função SetupRoutes de routes.go

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
