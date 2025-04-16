package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using default environment variables")
    }

	port := os.Getenv("PORT")
    if (port == "") {
        port = "8080"
    }


	router := gin.Default() //gin router instance

	router.GET("/", func(c *gin.Context) {

        c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"success": true,
            "message": "Hello, World!",
        })
    })
	
	if err := router.Run(": " + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }	
}