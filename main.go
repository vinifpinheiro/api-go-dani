package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// book represents data about a book.
type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// books slice to seed book data.
var books = []book{
	{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Price: 10.99},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 12.99},
	{ID: "3", Title: "1984", Author: "George Orwell", Price: 15.99},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:8080")
}

// getBooks responds with the list of all books as JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBook adds a book from JSON received in the request body.
func postBook(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// updateBook updates the details of a specific book.
func updateBook(c *gin.Context) {
	id := c.Param("id")

	var updatedBook book

	// Call BindJSON to bind the received JSON to
	// updatedBook.
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	// Find the book with the given ID and update its details.
	for i, b := range books {
		if b.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}

	// If no book with the given ID is found, return 404.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// deleteBook removes a specific book from the slice.
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	// Find the index of the book with the given ID and remove it from the slice.
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	// If no book with the given ID is found, return 404.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
