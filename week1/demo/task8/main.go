package main

import (
	"fmt"
	"sort"
)

func main() {
	firstSlice := []int{5, 3, 8, 2, 9, 4, 7, 1, 6}
	evenSlice := getEvenNumbers(firstSlice)
	fmt.Println("исходный слайс: ", firstSlice)
	fmt.Println("Только чётные: ", evenSlice)
	sort.Ints(evenSlice)
	fmt.Println("Отсортированные чётные: ", evenSlice)
}

func getEvenNumbers(slice []int) []int {
	var result []int
	for i := range slice {
		if slice[i]%2 == 0 {
			result = append(result, slice[i])
		}
	}
	return result
}
