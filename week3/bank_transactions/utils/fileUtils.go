package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %v\n", err)
		os.Exit(1)
	}
	return file
}

func GetNotEmptyFileLines(file *os.File) []string {
	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			result = append(result, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ошибка чтения: %v\n", err)
		os.Exit(1)
	}
	return result
}
