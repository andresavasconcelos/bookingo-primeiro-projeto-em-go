package models

type Review struct {
	Rating  int    `json:"rating"` // de 1 a 5
	Comment string `json:"comment"`
}
