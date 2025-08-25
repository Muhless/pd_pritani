package routes

import (
	"pd_pritani/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProfileRoutes(r *gin.Engine) {
	ProfileGroup := r.Group("/profiles")
	{
		ProfileGroup.POST("", handler.CreateProfile)
		ProfileGroup.GET("", handler.GetProfile)
		ProfileGroup.GET("/:id", handler.GetProfileByID)
		ProfileGroup.PATCH("/:id", handler.UpdateProfile)
		ProfileGroup.DELETE("/id", handler.DeleteProfile)
	}

}
