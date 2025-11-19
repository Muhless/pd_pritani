package routes

import (
	"pd_pritani/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(r *gin.Engine) {
	customerGroup := r.Group("/customers")
	{
		customerGroup.GET("", handler.GetCustomer)
		customerGroup.POST("", handler.CreateCustomer)
	}

}
