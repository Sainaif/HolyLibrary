package main

import (
	"log"

	"holyLibrary-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.InitDatabase()
	defer database.DB.Close()

	// Initialize Gin server
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to HolyLibrary API",
		})
	})

	log.Println("Starting server on :8080")
	r.Run(":8080") // Start the server on port 8080
}
