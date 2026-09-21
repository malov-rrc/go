package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var digits []int
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	left, right := 0, len(digits)-1
	for left < right {
		if digits[left] == digits[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(123))
}
