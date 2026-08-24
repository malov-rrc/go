package models

import "math"

type Circle struct {
	Radius int
}

func (c Circle) GetArea() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2.0)
}

func (c Circle) GetPerimeter() float64 {
	return 2 * math.Pi * float64(c.Radius)
}
