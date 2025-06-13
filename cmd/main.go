package main

import (
	"bookingo/internal/data"
	"bookingo/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadBooks()
	router := gin.Default()
	router.GET("/books", handler.GetBook)
	router.GET("/books/:id", handler.GetBookByID)
	router.POST("/books", handler.PostBook)
	router.DELETE("/books/:id", handler.RemoveBook)
	router.PUT("books/:id", handler.UpdateBookByID)
	router.POST("books/:id/reviews", handler.PostReview)
	router.Run()
}
