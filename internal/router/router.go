package router

import (
	"pd_pritani/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}
	return r
}
