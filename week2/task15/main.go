package main

import (
	"errors"
	"fmt"
)

// **15. Рекурсия**
// `func Factorial(n int) (int, error)`. `n < 0` → ошибка (не паника). `n == 0` → `1, nil`.
// Table-driven тест: `0→1`, `1→1`, `5→120`, `-1→error`.

func main() {
	fmt.Print(Factorial(3))
}

func Factorial(n int) (int, error) {
	switch {
	case n < 0:
		return 0, errors.New("Can't make factorial for number less than 0")
	case n == 0:
		return 1, nil
	default:
		x, _ := Factorial(n - 1)
		return x * n, nil
	}
}
