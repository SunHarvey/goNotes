package main

import "fmt"

type Gallons float64
type Liters float64
type Milliters float64

func main() {
	fmt.Println(Gallons(12.0935353))
	fmt.Println(Liters(12.0935353))
	fmt.Println(Milliters(12.0935353))

	fmt.Printf("%0.2f gal\n",Gallons(12.0935353))
	fmt.Printf("%0.2f L\n",Liters(12.0935353))
	fmt.Printf("%0.2f mL\n",Milliters(12.0935353))
}
