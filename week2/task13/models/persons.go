package models

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func SortPersons(persons []Person) {
	sort.Slice(persons, func(i, j int) bool {
		if persons[i].Age == persons[j].Age {
			return persons[i].Name < persons[j].Name
		}
		return persons[i].Age > persons[j].Age
	})
}

func PrintPersons(persons []Person) {
	for i := range persons {
		fmt.Printf("%s (%d)\n", persons[i].Name, persons[i].Age)
	}
}
