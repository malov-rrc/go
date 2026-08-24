package main

import "fmt"

func main() {
	counter1 := NewCounter()
	counter2 := NewCounter()
	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2
	fmt.Println(counter1()) // 2
	fmt.Println(counter2()) // 2
}

func NewCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
