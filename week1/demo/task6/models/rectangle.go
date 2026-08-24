package models

type Rectangle struct {
	Width  int
	Length int
}

func (r Rectangle) GetArea() float64 {
	return float64(r.Width) * float64(r.Length)
}

func (r Rectangle) GetPerimeter() float64 {
	return float64((2 * r.Length) + (2 * r.Width))
}
