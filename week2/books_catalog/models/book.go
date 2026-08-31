package models

import (
	"fmt"
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

func ParseLine(line string) (Book, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 4 {
		return Book{}, fmt.Errorf("invalid line format: expected 4 fields, got %d", len(parts))
	}

	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return Book{}, fmt.Errorf("invalid year: %w", err)
	}

	status := parts[3]
	switch status {
	case "read", "reading", "planned":
		// valid status
	default:
		status = "unknown"
	}

	return Book{
		Title:  parts[0],
		Author: parts[1],
		Year:   year,
		Status: status,
	}, nil
}

func SortBooks(books []Book) {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Year != books[j].Year {
			return books[i].Year < books[j].Year
		}
		return books[i].Title < books[j].Title
	})
}

func FormatCatalog(books []Book, invalidCount int) string {
	var out strings.Builder

	fmt.Fprintf(&out, "Total: %d\n", len(books))

	counts := make(map[string]int)
	for _, b := range books {
		counts[b.Status]++
	}

	var statusParts []string
	for _, status := range []string{"read", "reading", "planned", "unknown"} {
		if count := counts[status]; count > 0 {
			statusParts = append(statusParts, fmt.Sprintf("%s=%d", status, count))
		}
	}
	if len(statusParts) > 0 {
		fmt.Fprintf(&out, "By status: %s\n", strings.Join(statusParts, ", "))
	} else {
		fmt.Fprintln(&out, "By status:")
	}

	if invalidCount > 0 {
		fmt.Fprintf(&out, "Invalid lines: %d\n", invalidCount)
	}

	SortBooks(books)

	fmt.Fprintln(&out, "Sorted by year:")
	for _, b := range books {
		fmt.Fprintf(&out, "%d - %s (%s)\n", b.Year, b.Title, b.Author)
	}

	return out.String()
}
