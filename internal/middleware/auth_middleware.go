package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. ambil token dari header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token not found",
			})
			ctx.Abort()
			return
		}

		// 2. cek format bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Format token not valid",
			})
			ctx.Abort()
			return
		}

		// 3. validasi token
		tokenSting := parts[1]
		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenSting, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token doesn't valid",
			})
			ctx.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		ctx.Set("user_id", claims["user_id"])
		ctx.Set("role", claims["role"])

		ctx.Next()
	}
}
