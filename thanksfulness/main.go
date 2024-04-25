package main

import (
	"log"
	"my-tools/thanksfulness/controllers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/thanksfulness", controllers.Thanksfulness)
	r.Run() // http://localhost:8080
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			log.Println("No .env file")
		} else {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}
