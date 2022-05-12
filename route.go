package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// *gin.Context contains all of the information about the request
func getBooks(c *gin.Context) {
	books, err := getBooksFromDB()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Fail to retrieve list of books"})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id") // path parameter
	book, err := getBookByIdFromDB(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook Book

	// if there is an error binding the request json to newBook, c.BindJSON will return a bad request
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	_, err := createBookInDB(newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Book not created."})
		return
	}

	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id") // query parameter

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookByIdFromDB(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not avaliable."})
		return
	}

	book.Quantity -= 1
	_, err = editBookInDB(*book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Fail to checkout."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookByIdFromDB(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	book.Quantity += 1
	_, err = editBookInDB(*book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Fail to return."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}
