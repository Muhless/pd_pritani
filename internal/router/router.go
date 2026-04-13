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
	salesHandler *handler.SalesHandler) *gin.Engine {

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", authHandler.GetProfile)
	protected.PATCH("/profile", authHandler.UpdateProfile)

	{
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleGuard("admin"))
		{
			admin.POST("/register", authHandler.Register)
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

	// for check route
	// for _, route := range r.Routes() {
	// 	log.Println(route.Method, route.Path)
	// }

	return r
}
