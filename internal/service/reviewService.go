package service

import (
	"bookingo/internal/models"
	"errors"
)

func ValidateReviewRating(review models.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("A avaliação precisa ser de 1 a 5")
	}
	return nil
}
