package main

import (
	"github.com/gin-gonic/gin"
)

// Run the command: go run *.go
func main() {
	// gin router setup
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
