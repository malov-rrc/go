package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Проверяем, что передан аргумент
	if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <год>")
		os.Exit(1)
	}

	// Читаем первый аргумент (индекс 1)
	arg := os.Args[1]

	// Преобразуем строку в число (int)
	year, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Ошибка: '%s' не является корректным годом\n", arg)
		os.Exit(1)
	}

	// Вызываем функцию и выводим результат
	fmt.Printf("Год %d високосный: %t\n", year, isLeapYear(int16(year)))
}

func isLeapYear(year int16) bool {
	// Более правильная проверка на високосный год
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
