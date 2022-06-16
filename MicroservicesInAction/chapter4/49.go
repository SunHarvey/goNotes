package main

import (
	"fmt"
	"image/color"
)

type Rectangle struct{ w, h float64 }

type ColorRect struct {
	Rectangle
	Color color.RGBA
}

func main() {
	var cr ColorRect
	cr.h = 3
	cr.w = 2
	fmt.Println(cr.area())
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}
