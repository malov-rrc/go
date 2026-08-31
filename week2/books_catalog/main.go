package main

import (
	"books_catalog/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book = models.Book

func parseLine(line string) (Book, error) {
	return models.ParseLine(line)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: file path is required")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var books []Book
	invalidCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		book, err := parseLine(line)
		if err != nil {
			invalidCount++
			continue
		}
		books = append(books, book)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(models.FormatCatalog(books, invalidCount))
}
