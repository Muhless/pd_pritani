package routes

import (
	"pd_pritani/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(r *gin.RouterGroup) {
	customerGroup := r.Group("/customers")
	{
		customerGroup.GET("", handler.GetCustomers)
		customerGroup.POST("", handler.CreateCustomer)
	}

}
