package main

import (
	"log"
	"pd_pritani/internal/config"
	"pd_pritani/internal/routes"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// routes
	routes.RegisterUserRoutes(r)
	routes.RegisterAuthRoutes(r)
	routes.RegisterProductRoutes(r)
	routes.RegisterEmployeeRoutes(r)
	routes.RegisterCustomerRoutes(r)

	r.Static("/uploads", "./uploads")

	// port
	for _, ri := range r.Routes() {
		println(ri.Method, ri.Path)
	}

	r.Run(":8080")
}
