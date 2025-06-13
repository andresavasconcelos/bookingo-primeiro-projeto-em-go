package data

import (
	"bookingo/internal/models"
	"encoding/json"
	"fmt"
	"os"
)

var Books []models.Book

func LoadBooks() {
	file, err := os.Open("dados/books.json")
	if err != nil {
		fmt.Println("Error file: ", err)
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Books); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func SaveBook() {
	file, err := os.Create("dados/books.json")
	if err != nil {
		fmt.Println("Erroe file: ", err)
		return
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Books); err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}
}
