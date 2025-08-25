package main

import (
	"log"
	"pd_pritani/internal/config"
	"pd_pritani/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed connecting to .env")
	}

	config.ConnectDB()
	r := gin.Default()

	// routes
	routes.RegisterUserRoutes(r)
	routes.RegisterAuthRoutes(r)
	routes.RegisterProductRoutes(r)
	routes.RegisterProfileRoutes(r)
	r.Static("/uploads", "./uploads")
	// port
	r.Run(":8080")
}
