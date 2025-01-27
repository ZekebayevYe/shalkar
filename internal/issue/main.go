package main

import (
	"log"

	"github.com/NameSurname/Assignment1/GOLANGFINAL/config"
	"github.com/NameSurname/Assignment1/GOLANGFINAL/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	router := gin.Default()
	routes.RegisterRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
