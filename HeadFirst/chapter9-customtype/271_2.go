package main

import "fmt"

type Liters float64
type Gallons float64

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.246)
}

func ToLiters(l Gallons) Liters {
	return Liters(l * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}
