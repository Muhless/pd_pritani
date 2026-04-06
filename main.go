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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed connecting to .env")
	}

	db := config.ConnectDB()

	// user
	userRepo := repository.NewUserRepository(db)
	adminRepo := repository.NewAdminRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)

	authService := service.NewAuthService(db, userRepo, adminRepo, employeeRepo)
	authHandler := handler.NewAuthHandler(authService)

	// customer
	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	// products
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r := router.SetupRouter(authHandler, customerHandler, productHandler)

	r.Run(":8080")
}
