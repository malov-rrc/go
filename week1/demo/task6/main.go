package main

import (
	"fmt"
	"task6/models"
)

func main() {
	var f models.Figure
	f = models.Rectangle{Width: 3, Length: 5}
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", f.GetArea(), f.GetPerimeter())
	f = models.Circle{Radius: 5}
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", f.GetArea(), f.GetPerimeter())
}
