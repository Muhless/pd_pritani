package routes

import (
	"pd_pritani/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProfileRoutes(r *gin.Engine) {
	ProfileGroup := r.Group("/profiles")
	{
		ProfileGroup.POST("", handler.CreateEmployee)
		ProfileGroup.GET("", handler.GetEmployee)
		ProfileGroup.GET("/:id", handler.GetEmployeeByID)
		ProfileGroup.PATCH("/:id", handler.UpdateEmployee)
		ProfileGroup.DELETE("/id", handler.DeleteEmployee)
	}

}
