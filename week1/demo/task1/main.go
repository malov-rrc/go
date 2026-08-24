package main

import (
	"fmt"
)

func main() {
	fmt.Println("Задача 1. 32 градусов цельсия в фаренгейт: ", celsiusToFahrenheit(32))
	fmt.Println("Задача 1. -14 градусов фаренгейта в цельсий: ", fahrenheitToCelsius(-14))
}

// задача 1 - просто познакомиться с типами и переменными, поэтому просто создам 2 разные функции, без геморроя
func celsiusToFahrenheit(celsius float32) float32 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(farenheit float32) float32 {
	return (farenheit - 32) * 5 / 9
}
