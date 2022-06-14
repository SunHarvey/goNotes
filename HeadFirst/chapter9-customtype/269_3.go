package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.2)
	fmt.Println(Gallons(1.2) == 1.2)
	fmt.Println(Liters(1.2) < 3.4)
}
