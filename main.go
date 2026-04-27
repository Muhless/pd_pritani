// @title 			PG Pritani API
// version 			1.0
// @description			API for PG Pritani dashboard
// @host				localhost:8080
// @BasePath			/
// @securityDefinitions.apikey 	BearerAuth
// @in Header
// @name Authorization
package main

import (
	"pd_pritani/internal/config"
	"pd_pritani/internal/handler"
	"pd_pritani/internal/repository"
	"pd_pritani/internal/router"
	"pd_pritani/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

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

	// sales
	salesRepo := repository.NewSalesRepository(db)
	salesService := service.NewSalesService(db, salesRepo, productRepo, employeeRepo)
	salesHandler := handler.NewSalesHandler(salesService)

	// supplier
	supplierRepo := repository.NewSupplierRepository(db)
	supplierService := service.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	// purchase
	purchaseRepo := repository.NewPurchaseRepository(db)
	purchaseService := service.NewPurchaseService(purchaseRepo)
	purchaseHandler := handler.NewPurchaseHandler(purchaseService)

	r := router.SetupRouter(authHandler, customerHandler, productHandler, salesHandler, supplierHandler, purchaseHandler)

	r.Run(":8080")
}
