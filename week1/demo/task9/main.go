package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	text := "task9/main.go  - Частота слов в строке (map). " +
		"Источник данных - захардкоженная строка-константа (2–3 предложения) в main." +
		"Вывод — построчно, отсортировано по алфавиту: слово: количество."
	mapWithWordCount := wordCounter(text)
	sortedKeys := sortKeys(mapWithWordCount)
	for i := range sortedKeys {
		fmt.Printf("%s: %d\n", sortedKeys[i], mapWithWordCount[sortedKeys[i]])
	}
}

func sortKeys(mapWithWordCount map[string]int) []string {
	keys := make([]string, 0, len(mapWithWordCount))
	for k := range mapWithWordCount {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func wordCounter(text string) map[string]int {
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	result := make(map[string]int)

	for _, w := range words {
		lower := strings.ToLower(w)
		result[lower]++
	}

	return result
}
