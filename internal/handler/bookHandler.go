package handler

import (
	"bookingo/internal/data"
	"bookingo/internal/models"
	"bookingo/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoveBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}

	for i, book := range data.Books {
		if book.ID == id {
			data.Books = append(data.Books[:i], data.Books[i+1:]...)
			data.SaveBook()
			c.JSON(200, gin.H{"message": "livro removido"})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "livro não encontrado"})
}

func GetBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"books": data.Books,
	})
}

func GetBookByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	for _, book := range data.Books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "livro não encontrado"})
}

func PostBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	if err := service.ValidateBookPrice(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
	}

	newBook.ID = len(data.Books) + 1
	data.Books = append(data.Books, newBook)
	data.SaveBook()
	c.JSON(http.StatusCreated, newBook)
}

func putBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"books": data.Books,
	})
}

func UpdateBookByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	var updated models.Book
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidateBookPrice(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
	}

	for i, book := range data.Books {
		if book.ID == id {
			data.Books[i] = updated
			data.Books[i].ID = id
			data.SaveBook()
			c.JSON(200, data.Books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"method": "livro não encontrado"})
}
