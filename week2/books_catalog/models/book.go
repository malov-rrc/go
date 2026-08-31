package models

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Status string
}

func parseLine(line string) (Book, error) {
	splittedLine := strings.Split(line, ",")
	if len(splittedLine) != 4 {
		return Book{}, errors.New("invalid Line")
	}
	year, err := strconv.Atoi(splittedLine[2])
	if err != nil {
		year = -1
	}
	status := splittedLine[3]
	switch status {
	case "read", "reading", "planned":
		//ok
	default:
		status = "unknown"
	}
	return Book{
		Title:  splittedLine[0],
		Author: splittedLine[1],
		Year:   year,
		Status: status,
	}, nil
}

func SortBooksByYear(books []Book) {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Year == books[j].Year {
			return books[i].Title < books[j].Title
		}
		return books[i].Year < books[j].Year
	})
}

func GetBooksFromLines(lines []string) ([]Book, []error) {
	var books []Book
	var errorList []error
	for i := range lines {
		parsedBook, err := parseLine(lines[i])
		if err != nil {
			errorList = append(errorList, err)
		} else {
			books = append(books, parsedBook)
		}
	}
	return books, errorList
}
