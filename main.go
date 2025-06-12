package main

import (
	"bookingo/models"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books []models.Book

func main() {
	fmt.Println("Hello World!")
}

func loadBooks() {
	file, err := os.Open("dados/books.json")
	if err != nil {
		fmt.Println("Error file: ", err)
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&books); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func saveBook() {
	file, err := os.Create("dados/books.json")
	if err != nil {
		fmt.Println("Erroe file: ", err)
		return
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(books); err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}
}

func removeBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			saveBook()
			c.JSON(200, gin.H{"message": "book deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "book not found"})
}

func getBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"books": books,
	})
}

func getBookByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	for _, book := range books {
		if book.ID == id {
			c.JSON(200, book)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Book not found"})
}

func postBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	saveBook()
	c.JSON(201, newBook)
}

func putBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"books": books,
	})
}

func updateBookByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	var updated models.Book
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updated
			books[i].ID = id
			saveBook()
			c.JSON(200, books[i])
			return
		}
	}

	c.JSON(404, gin.H{"method": "book not found"})
}
