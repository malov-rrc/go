package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Использование: go run main.go <N>, где N - оценка")
		os.Exit(1)
	} else if len(os.Args) < 2 {
		fmt.Println("Использование: go run main.go <N>, где N - оценка")
		os.Exit(1)
	} else {
		var grade uint64 = 2
		arg := os.Args[1]
		grade, err := strconv.ParseUint(arg, 10, 8)
		if err != nil {
			fmt.Printf("Ошибка: '%s' не является корректным числом\n", arg)
			os.Exit(1)
		}
		fmt.Println("Ваша оценка:", gradeClassification(grade))
	}

}

func gradeClassification(grade uint64) string {
	result := ""
	switch grade {
	case 5:
		result = "Отлично"
	case 4:
		result = "Хорошо"
	case 3:
		result = "Удовлетворительно"
	case 2:
		result = "Неудовлетворительно"
	case 1:
		result = "Плохо"
	default:
		os.Exit(1)
	}
	return result
}
