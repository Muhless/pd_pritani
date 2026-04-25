package router

import (
	"pd_pritani/internal/handler"
	"pd_pritani/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authHandler *handler.AuthHandler,
	customerHandler *handler.CustomerHandler,
	productHandler *handler.ProductHandler,
	salesHandler *handler.SalesHandler,
	supplierHandler *handler.SupplierHandler,
	purchaseHandler *handler.PurchaseHandler,
) *gin.Engine {

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", authHandler.GetProfile)
	protected.PATCH("/profile", authHandler.UpdateProfile)

	{
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleGuard("admin"))
		{
			admin.GET("/users", authHandler.GetAllUsers)
			admin.GET("/users/:id", authHandler.GetUserByID)
			admin.PATCH("/users/:id", authHandler.UpdateUser)
			admin.DELETE("/users/:id", authHandler.DeleteUser)
		}
	}

	customer := protected.Group("/customers")
	{
		customer.GET("/", customerHandler.GetAll)
		customer.GET("/:id", customerHandler.GetByID)
		customer.POST("/", customerHandler.Create)
		customer.PATCH("/:id", customerHandler.Update)
		customer.DELETE("/:id", customerHandler.Delete)
	}

	products := protected.Group("/products")
	{
		products.GET("/", productHandler.GetAll)
		products.GET("/:id", productHandler.GetByID)
		products.POST("/", productHandler.Create)
		products.PATCH("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)
	}

	sales := protected.Group("/sales")
	{
		sales.GET("/", salesHandler.GetAll)
		sales.GET("/:id", salesHandler.GetByID)
		sales.POST("/", salesHandler.Create)
		sales.PATCH("/:id", salesHandler.UpdateStatus)
		sales.DELETE("/:id", salesHandler.Delete)
	}

	supplier := protected.Group("/suppliers")
	{
		supplier.GET("/", supplierHandler.GetAll)
		supplier.GET("/:id", supplierHandler.GetByID)
		supplier.POST("/", supplierHandler.Create)
		supplier.PATCH("/:id", supplierHandler.Update)
		supplier.DELETE("/:id", supplierHandler.Delete)
	}
	purchase := protected.Group("/purchases")
	{
		purchase.GET("/", purchaseHandler.GetAll)
		purchase.GET("/:id", purchaseHandler.GetByID)
		purchase.POST("/", purchaseHandler.Create)
		purchase.PATCH("/:id", purchaseHandler.Update)
		purchase.DELETE("/:id", purchaseHandler.Delete)
	}

	// for check route
	// for _, route := range r.Routes() {
	// 	log.Println(route.Method, route.Path)
	// }

	return r
}
