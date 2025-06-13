package handler

import (
	"bookingo/internal/data"
	"bookingo/internal/models"
	"bookingo/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var newReview models.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidateReviewRating(newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, book := range data.Books {
		if book.ID == bookID {
			book.Review = append(book.Review, newReview)
			data.Books[i] = book
			data.SaveBook()
			c.JSON(http.StatusCreated, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book não encontrado"})
}
