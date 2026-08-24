package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "укажите путь к файлу")
		os.Exit(1)
	}
	filename := os.Args[1]
	file := openFile(filename)
	defer file.Close()
	fileLines := getNotEmptyFileLines(file)
	testResults := getTestResultsFromText(fileLines)
	fmt.Print(getTestStat(testResults))
}

type TestResult struct {
	name   string
	status string
}

func getTestStat(testResults []TestResult) string {
	total := len(testResults)
	counts := make(map[string]int)
	for _, r := range testResults {
		counts[r.status]++
	}

	out := fmt.Sprintf("Total: %d\n", total)
	for _, status := range []string{"PASS", "FAIL", "SKIP", "unknownStatus"} {
		if count, ok := counts[status]; ok && count > 0 {
			out += fmt.Sprintf("%s: %d (%.1f%%)\n", status, count, float64(count)/float64(total)*100)
		}
	}
	if count, ok := counts["invalidLine"]; ok && count > 0 {
		out += fmt.Sprintf("Invalid lines: %d\n", count)
	}
	return out
}

func getTestResultsFromText(textLines []string) []TestResult {
	var testResults []TestResult
	for i := range textLines {
		splittedLine := strings.Split(textLines[i], ",")
		if len(splittedLine) != 2 {
			testResults = append(testResults, TestResult{"unknownName(invalid result line)", "invalidLine"})
			continue
		}
		testName := splittedLine[0]
		testStatus := splittedLine[1]
		switch testStatus {
		case "PASS", "FAIL", "SKIP":
			testResults = append(testResults, TestResult{testName, testStatus})
		default:
			testResults = append(testResults, TestResult{testName, "unknownStatus"})
		}
	}
	return testResults
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %v\n", err)
		os.Exit(1)
	}
	return file
}

func getNotEmptyFileLines(file *os.File) []string {
	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
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
