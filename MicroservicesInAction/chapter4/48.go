package main

import "fmt"

type Rectangle struct{ w, h float64 }

func main() {
	p := &Rectangle{w: 2, h: 3}
	fmt.Println(p.area())

	rec := Rectangle{w: 2, h: 3}
	rp := &rec
	fmt.Println(rp.area())

	fmt.Println((&rec).area())
	fmt.Println(rec.area())

	//Rectangle{w:2, h:3}.area()
	Rectangle{w: 2, h: 3}.area2()
}

func (r *Rectangle) area() float64 {
	return r.w * r.h
}

func (r Rectangle) area2() float64 {
	return r.w * r.h
}
