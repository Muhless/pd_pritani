package routes

import (
	"pd_pritani/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/products")
	{
		productGroup.GET("/", handler.GetProduct)
		productGroup.GET("/:id", handler.GetProductByID)
		productGroup.POST("/", handler.CreateProduct)
		productGroup.PATCH("/:id", handler.UpdateProduct)
		productGroup.DELETE("/:id", handler.DeleteProduct)
	}

}
