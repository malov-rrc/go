package main

import (
	"errors"
)

func main() {
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на 0 невозможно (пока что)")
	} else {
		return a / b, nil
	}
}
