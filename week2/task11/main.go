package main

import "fmt"

func main() {
	fmt.Println("SumAll(1, 2, 3) =", SumAll(1, 2, 3))
	fmt.Println("SumAll() =", SumAll())
	fmt.Println("SumAll(5) =", SumAll(5))
}

func SumAll(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
