package utils

import (
	"io"
	"os"
	"slices"
	"testing"
)

func TestGetNotEmptyFileLines(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result []string
	}{
		{"empty string", "    ", []string{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test-*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer func(name string) {
				err := os.Remove(name)
				if err != nil {
					t.Fatal(err)
				}
			}(tmpFile.Name())
			defer func(tmpFile *os.File) {
				err := tmpFile.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(tmpFile)

			if _, err := tmpFile.WriteString(tc.input); err != nil {
				t.Fatal(err)
			}

			if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
				t.Fatal(err)
			}

			actualResult := GetNotEmptyFileLines(tmpFile)
			if !slices.Equal(actualResult, tc.result) {
				t.Errorf("GetNotEmptyFileLines with file. that has string \"%s\" result invalid. Expected: %q, got %q", tc.input, tc.result, actualResult)
			}
		})
	}
}
