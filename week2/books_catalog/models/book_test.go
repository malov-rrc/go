package models

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue Book
		wantErr   bool
	}{
		{"valid string", "1984,George Orwell,1949,read", Book{Title: "1984", Author: "George Orwell", Year: 1949, Status: "read"}, false},
		{"field lines more than expected", "1984,George Orwell,1949,read,1", Book{}, true},
		{"field lines less than expected", "1984,George Orwell,1949", Book{}, true},
		{"invalid year", "1984,George Orwell,ffff,read", Book{Title: "1984", Author: "George Orwell", Year: -1, Status: "read"}, false},
		{"empty string", "", Book{}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseLine(tc.input)
			if tc.wantErr && err == nil {
				t.Error("Error is expected, but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error, got %v", err)
			}
			if got != tc.wantValue {
				t.Errorf("parseLine(\"%s\") = %v, want %v", tc.input, got, tc.wantValue)
			}
		})
	}
}
