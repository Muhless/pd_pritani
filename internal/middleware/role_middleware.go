package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleGuard(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "Role doesn't found",
			})
			ctx.Abort()
			return
		}

		for _, r := range roles {
			if r == role.(string) {
				ctx.Next()
				return
			}
		}
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "Didn't have permission",
		})
		ctx.Abort()
	}

}
