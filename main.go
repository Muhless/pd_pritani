package main

import (
	"log"
	"pd_pritani/internal/config"
	"pd_pritani/internal/handler"
	"pd_pritani/internal/repository"
	"pd_pritani/internal/router"
	"pd_pritani/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed connecting to .env")
	}

	db := config.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := router.SetupRouter(authHandler)

	r.Run(":8080")
}
