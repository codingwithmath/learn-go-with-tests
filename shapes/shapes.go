package shapes

import "math"

// In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Heigth float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Heigth float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Heigth)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Heigth
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Heigth
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Heigth) * 0.5
}
