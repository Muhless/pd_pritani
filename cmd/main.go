package main

import (
	"log"
	"pd_pritani/api/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed connecting to .env")
	}

	config.ConnectDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome"})
	})

	r.Run(":8080")
}
