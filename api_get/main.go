package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct to hold the message
type Response struct {
	Value int `json:"Value"`
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a GET endpoint
	r.GET("/one", func(c *gin.Context) {
		msg := Response{Value: 1}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/two", func(c *gin.Context) {
		msg := Response{Value: 2}
		c.JSON(http.StatusOK, msg)
	})

	// Start the server on port 8080
	r.Run(":8080")
}
