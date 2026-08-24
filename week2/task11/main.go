package main

import "fmt"

//**11. Вариативные функции**
// `func SumAll(nums ...int) int`. В `main` вызвать и напечатать:
// `SumAll(1, 2, 3) = 6`, `SumAll() = 0`, `SumAll(5) = 5`.

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
