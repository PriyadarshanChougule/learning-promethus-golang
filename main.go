package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func check() {
	// Create a new Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Example: root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin server!",
		})
	})

	// Example: simple API route
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "pong",
		})
	})

	// Example: POST route
	router.POST("/api/data", func(c *gin.Context) {
		var requestBody map[string]interface{}
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"received": requestBody,
		})
	})

	// Start server on port 8080
	router.Run(":8080")
}
