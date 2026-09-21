package main

import (
	"fmt"
	"task13/models"
)

func main() {
	persons := []models.Person{
		{Name: "Anna", Age: 34},
		{Name: "Boris", Age: 22},
		{Name: "Vera", Age: 45},
		{Name: "Gleb", Age: 22},
	}
	fmt.Println("--------------------- Before sort: ---------------------")
	models.PrintPersons(persons)
	models.SortPersons(persons)
	fmt.Println("--------------------- After sort: ---------------------")
	models.PrintPersons(persons)
}
