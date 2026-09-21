package main

import (
	"fmt"
	"maps"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapper1 := make(map[rune]int)
	mapper2 := make(map[rune]int)
	for _, v := range s {
		mapper1[v]++
	}
	for _, v := range t {
		mapper2[v]++
	}
	return maps.Equal(mapper1, mapper2)
}

func main() {
	fmt.Println(isAnagram("aba", "bab"))
	fmt.Println(isAnagram("aba", "baa"))
}
