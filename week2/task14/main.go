package main

import "fmt"

type Counter struct{ count int }

// Нужен именно pointer receiver, т.к. при получении value - мы будем увеличивать переменную не у полученного
// объекта, а у его копии внутри метода (В го наверно это не объъектами называется как в джаве, но я хз).
// Вроде так по памяти
func (c *Counter) Increment() {
	c.count++
}

func main() {
	counter := Counter{}
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println(counter)
}
