package main

import (
	"books_catalog/models"
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantBook Book
		wantErr  bool
	}{
		{
			name:  "valid line with read status",
			input: "1984,George Orwell,1949,read",
			wantBook: Book{
				Title:  "1984",
				Author: "George Orwell",
				Year:   1949,
				Status: "read",
			},
			wantErr: false,
		},
		{
			name:  "valid line with reading status",
			input: "Dune,Frank Herbert,1965,reading",
			wantBook: Book{
				Title:  "Dune",
				Author: "Frank Herbert",
				Year:   1965,
				Status: "reading",
			},
			wantErr: false,
		},
		{
			name:  "valid line with planned status",
			input: "Neuromancer,William Gibson,1984,planned",
			wantBook: Book{
				Title:  "Neuromancer",
				Author: "William Gibson",
				Year:   1984,
				Status: "planned",
			},
			wantErr: false,
		},
		{
			name:  "valid line with unknown status",
			input: "Neuromancer,William Gibson,1984,lol",
			wantBook: Book{
				Title:  "Neuromancer",
				Author: "William Gibson",
				Year:   1984,
				Status: "unknown",
			},
			wantErr: false,
		},
		{
			name:     "invalid line with too few fields",
			input:    "Neuromancer,William Gibson,1984",
			wantBook: Book{},
			wantErr:  true,
		},
		{
			name:     "invalid line with too many fields",
			input:    "Title,Author,1990,read,extra",
			wantBook: Book{},
			wantErr:  true,
		},
		{
			name:     "invalid line with non-numeric year",
			input:    "Title,Author,year_text,read",
			wantBook: Book{},
			wantErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseLine(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("parseLine(%q) expected error, got nil", tc.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseLine(%q) unexpected error: %v", tc.input, err)
			}
			if !reflect.DeepEqual(got, tc.wantBook) {
				t.Errorf("parseLine(%q) = %+v, want %+v", tc.input, got, tc.wantBook)
			}
		})
	}
}

func TestFormatCatalog(t *testing.T) {
	tests := []struct {
		name         string
		books        []Book
		invalidCount int
		wantOutput   string
	}{
		{
			name: "all valid books with standard statuses",
			books: []Book{
				{Title: "1984", Author: "George Orwell", Year: 1949, Status: "read"},
				{Title: "Dune", Author: "Frank Herbert", Year: 1965, Status: "reading"},
				{Title: "Neuromancer", Author: "William Gibson", Year: 1984, Status: "planned"},
			},
			invalidCount: 0,
			wantOutput: "Total: 3\n" +
				"By status: read=1, reading=1, planned=1\n" +
				"Sorted by year:\n" +
				"1949 - 1984 (George Orwell)\n" +
				"1965 - Dune (Frank Herbert)\n" +
				"1984 - Neuromancer (William Gibson)\n",
		},
		{
			name: "includes unknown status, invalid lines and duplicate years sorted by title",
			books: []Book{
				{Title: "1984", Author: "George Orwell", Year: 1949, Status: "read"},
				{Title: "Dune", Author: "Frank Herbert", Year: 1965, Status: "reading"},
				{Title: "Neuromancer", Author: "William Gibson", Year: 1984, Status: "planned"},
				{Title: "Count Zero", Author: "William Gibson", Year: 1984, Status: "unknown"},
			},
			invalidCount: 2,
			wantOutput: "Total: 4\n" +
				"By status: read=1, reading=1, planned=1, unknown=1\n" +
				"Invalid lines: 2\n" +
				"Sorted by year:\n" +
				"1949 - 1984 (George Orwell)\n" +
				"1965 - Dune (Frank Herbert)\n" +
				"1984 - Count Zero (William Gibson)\n" +
				"1984 - Neuromancer (William Gibson)\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := models.FormatCatalog(tc.books, tc.invalidCount)
			if got != tc.wantOutput {
				t.Errorf("FormatCatalog() mismatch:\nGOT:\n%s\nWANT:\n%s", got, tc.wantOutput)
			}
		})
	}
}
