package main

import (
	"books_catalog/models"
)

type Book = models.Book

func parseLine(line string) (Book, error) {
	return models.ParseLine(line)
}

func main() {

}
