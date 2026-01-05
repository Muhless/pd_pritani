package routes

import (
	"pd_pritani/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine)  {
	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/", handler.GetUsers)
		userGroup.GET("/:id", handler.GetUserByID)
		userGroup.POST("/", handler.CreateUser)
		userGroup.PATCH("/:id", handler.UpdateUser)
		userGroup.DELETE("/:id", handler.DeleteUser)
	}
}