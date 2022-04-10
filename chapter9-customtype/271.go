package main

//import "fmt"

type Liters float64
type Gallons float64

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(2.5)
	carFuel += Liters(8.0)
	busFuel += Gallons(30.0)
}
