package main

import (
	"errors"
	"math"
)

func main() {
}

func Sum(array []int) (int, error) {
	if len(array) < 1 {
		return 0, errors.New("Пустой массив")
	}
	var result = array[0]
	for i := 1; i < len(array); i++ {
		result += array[i]
	}
	return result, nil
}

func Max(array []int) (int, error) {
	if len(array) < 1 {
		return 0, errors.New("Пустой массив")
	}
	var max = array[0]
	for i := 1; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	return max, nil
}

func Min(array []int) (int, error) {
	if len(array) < 1 {
		return 0, errors.New("Пустой массив")
	}
	var min = array[0]
	for i := 1; i < len(array); i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	return min, nil
}

func IsPrime(testInt int) bool {
	if testInt <= 1 {
		return false
	} else if testInt%2 == 0 {
		return testInt == 2
	} else if testInt%3 == 0 {
		return testInt == 3
	} else {
		var limit = int(math.Sqrt(float64(testInt)))
		for i := 2; i <= limit; i++ {
			if testInt%i == 0 {
				return false
			}
		}
	}
	return true
}
