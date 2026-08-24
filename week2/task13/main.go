package main

import (
	"fmt"
	"sort"
)

// **13. Сортировка через `sort.Slice`**
// `type Person struct { Name string; Age int }`. Хардкод: `{"Anna",34}, {"Boris",22}, {"Vera",45}, {"Gleb",22}`.
// Сортировать по `Age` **по убыванию**; при равном `Age` — по `Name` **по возрастанию** (детерминированность).
// Напечатать список до и после сортировки, построчно: `Имя (Возраст)`.
type Person struct {
	Name string
	Age  int
}

func main() {
	persons := []Person{
		{"Anna", 34},
		{"Boris", 22},
		{"Vera", 45},
		{"Gleb", 22},
	}
	fmt.Println("Before sort: ")
	for i := range persons {
		fmt.Printf("%s (%d)\n", persons[i].Name, persons[i].Age)
	}
	sortPersons(persons)
	fmt.Println("After sort: ")
	for i := range persons {
		fmt.Printf("%s (%d)\n", persons[i].Name, persons[i].Age)
	}
}

func sortPersons(persons []Person) {
	sort.Slice(persons, func(i, j int) bool {
		if persons[i].Age == persons[j].Age {
			return persons[i].Name < persons[j].Name
		}
		return persons[i].Age > persons[j].Age
	})
}
