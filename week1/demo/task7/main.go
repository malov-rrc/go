package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 6
	fmt.Println("Before swap: a =", a, ", b =", b)
	Swap(&a, &b)
	fmt.Println("After swap: a =", a, ", b =", b)
}

func Swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}
