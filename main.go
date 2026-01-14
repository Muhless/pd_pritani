package main

import (
	"log"
	"pd_pritani/internal/config"
	"pd_pritani/internal/middleware"
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

	api := r.Group("/api")
	routes.RegisterAuthRoutes(api)

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		routes.RegisterUserRoutes(protected)
		routes.RegisterProductRoutes(protected)
		routes.RegisterEmployeeRoutes(protected)
		routes.RegisterCustomerRoutes(protected)
	}

	r.Static("/uploads", "./uploads")

	// port
	for _, ri := range r.Routes() {
		println(ri.Method, ri.Path)
	}

	r.Run(":8080")
}
