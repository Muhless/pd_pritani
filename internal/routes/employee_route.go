package routes

import (
	"pd_pritani/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(r *gin.Engine) {
	EmployeeGroup := r.Group("/employees")
	{
		EmployeeGroup.POST("", handler.CreateEmployee)
		EmployeeGroup.GET("", handler.GetEmployee)
		EmployeeGroup.GET("/:id", handler.GetEmployeeByID)
		EmployeeGroup.PATCH("/:id", handler.UpdateEmployee)
		EmployeeGroup.DELETE("/:id", handler.DeleteEmployee)
	}

}
