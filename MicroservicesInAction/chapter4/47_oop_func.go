package main

import "fmt"

type Rectangle struct{ w, h float64 }

func main() {
	rec := Rectangle{w: 2, h: 3}
	fmt.Println(area(rec.w, rec.h))
	fmt.Println(rec.area())
}

func area(w, h float64) float64 {
	return w * h
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}
