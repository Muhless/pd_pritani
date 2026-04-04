package router

import (
	"pd_pritani/internal/handler"
	"pd_pritani/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authHandler *handler.AuthHandler,
	customerHandler *handler.CustomerHandler) *gin.Engine {

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
		customer.PATCH("/", customerHandler.Update)
		customer.DELETE("/", customerHandler.Delete)
	}

	// for check route
	// for _, route := range r.Routes() {
	// 	log.Println(route.Method, route.Path)
	// }

	return r
}
