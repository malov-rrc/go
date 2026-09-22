package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	result := strs[0]
	for _, str := range strs[1:] {
		for result != "" && !strings.HasPrefix(str, result) {
			result = result[:len(result)-1]
		}
		if result == "" {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"b", "a", "aa"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
