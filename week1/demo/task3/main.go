package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Использование: go run main.go <N>, где N - размер таблицы NxN")
		os.Exit(1)
	} else if len(os.Args) < 1 {
		fmt.Println("Использование: go run main.go <год>")
		os.Exit(1)
	} else {
		var size uint64 = 10
		if len(os.Args) > 1 {
			parsed, err := strconv.ParseUint(os.Args[1], 10, 8)
			if err != nil {
				fmt.Printf("Ошибка: '%s' не является корректным числом\n", os.Args[1])
				os.Exit(1)
			}
			size = uint64(parsed)
		}

		// Вызываем функцию и выводим результат
		fmt.Printf("Матрица размера %dx%d: \n", size, size)
		multiplicationTable(size)
	}
}

func multiplicationTable(n uint64) {
	for i := uint64(1); i <= n; i++ {
		for j := uint64(1); j <= n; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Print("\n")
	}
}
