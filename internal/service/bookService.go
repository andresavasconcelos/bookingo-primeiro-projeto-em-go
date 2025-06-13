package service

import (
	"bookingo/internal/models"
	"errors"
)

// Valida se o valor do preço do livro é menos que zero, dado que o valor das coisas nao pode ser negativo.
func ValidateBookPrice(book *models.Book) error {
	if book.Preco < 0 {
		return errors.New("O preço do livro não pode ser negativo")
	}

	return nil
}
