package main

import (
	"books_catalog/models"
	"books_catalog/utils"
	"fmt"
	"os"
	"strings"
)

type Book = models.Book

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "укажите путь к файлу")
		os.Exit(1)
	}
	filename := os.Args[1]
	file := utils.OpenFile(filename)
	defer file.Close()
	fileLines := utils.GetNotEmptyFileLines(file)
	books, errors := models.GetBooksFromLines(fileLines)
	fmt.Print(getTestStat(books, errors))
}

func getTestStat(books []Book, errors []error) string {
	total := len(books) + len(errors)
	counts := make(map[string]int)
	for _, r := range books {
		counts[r.Status]++
	}
	var out strings.Builder
	fmt.Fprintf(&out, "Total: %d", total)
	fmt.Fprint(&out, "\nBy status: ")
	for index, status := range []string{"read", "reading", "planned", "unknown"} {
		if count, ok := counts[status]; ok && count > 0 {
			if index != 0 {
				fmt.Fprint(&out, ", ")
			}
			out.WriteString(fmt.Sprintf("%s: %d", status, count))
		}
	}
	if len(errors) > 0 {
		out.WriteString(fmt.Sprintf("\nInvalid lines: %d", len(errors)))
	}
	fmt.Fprint(&out, "\nSorted by year:\n")
	models.SortBooksByYear(books)
	for i := range books {
		out.WriteString(fmt.Sprintf("%d - %s\n", books[i].Year, books[i].Title))
	}
	return out.String()
}
