package main

import (
	"fmt"
	"math"
)

type ShapeDesc interface {
	Area() float64
	Perimeter() float64
}

type rectangle struct {
	H, W float64
}

type circle struct {
	R float64
}

func (r rectangle) Area() float64 {
	return r.H * r.W
}
func (r rectangle) Perimeter() float64 {
	return 2 * (r.H + r.W)
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}
func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func main() {
	var s1, s2 ShapeDesc
	s1 = rectangle{H: 2, W: 3}
	s2 = circle{R: 2}
	Desc(s1)
	Desc(s2)
}

func Desc(s ShapeDesc) {
	switch kind := s.(type) {
	case circle:
		fmt.Println("This is circle.")
	case rectangle:
		fmt.Println("This is rectangle.")
	default:
		fmt.Println("%v is unknown type", kind)
	}

	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
