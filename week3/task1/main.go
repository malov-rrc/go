package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CleanInput("   Hello    World  "))
}

func CleanInput(s string) string {
	s = strings.ToLower(s)
	return strings.Join(strings.Fields(s), " ")
}
