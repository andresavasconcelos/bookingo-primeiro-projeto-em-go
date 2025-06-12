package models

type Book struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Autor string  `json:"autor"`
	Ano   uint    `json:"uint"`
	Preco float64 `json:"preco"`
}
