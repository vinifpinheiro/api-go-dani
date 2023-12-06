package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Book representa dados sobre um livro.
type Book struct {
	ID     string  `json:"id" gorm:"primary_key"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("falha ao conectar ao banco de dados")
	}
	defer db.Close()

	// Migrar o esquema
	db.AutoMigrate(&Book{})

	router := gin.Default()
	router.GET("/livros", getBooks)
	router.POST("/livros", postBook)
	router.PUT("/livros/:id", updateBook)
	router.DELETE("/livros/:id", deleteBook)
	router.Run("localhost:8080")
}

// getBooks responde com a lista de todos os livros em JSON.
func getBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.IndentedJSON(http.StatusOK, books)
}

// postBook adiciona um livro a partir do JSON recebido no corpo da requisição.
func postBook(c *gin.Context) {
	var newBook Book

	// Chame BindJSON para vincular o JSON recebido a newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Adicione o novo livro ao banco de dados.
	db.Create(&newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// updateBook atualiza os detalhes de um livro específico.
func updateBook(c *gin.Context) {
	id := c.Param("id")

	var updatedBook Book

	// Chame BindJSON para vincular o JSON recebido a updatedBook.
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	// Encontre o livro com o ID fornecido e atualize seus detalhes.
	var book Book
	if db.First(&book, id).RecordNotFound() {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Livro não encontrado"})
		return
	}

	db.Model(&book).Updates(updatedBook)
	c.IndentedJSON(http.StatusOK, updatedBook)
}

// deleteBook remove um livro específico do banco de dados.
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	// Encontre o livro com o ID fornecido e exclua-o do banco de dados.
	var book Book
	if db.First(&book, id).RecordNotFound() {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Livro não encontrado"})
		return
	}

	db.Delete(&book)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Livro deletado"})
}
