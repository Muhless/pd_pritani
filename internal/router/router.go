package router

import (
	"log"
	"pd_pritani/internal/handler"
	"pd_pritani/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", authHandler.GetProfile)
	
	{
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleGuard("admin"))
		{
			admin.POST("/register", authHandler.Register)
		}
	}
	for _, route := range r.Routes() {
		log.Println(route.Method, route.Path)
	}
	return r
}
