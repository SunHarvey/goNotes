package main

import (
	"fmt"
	"geo"
)

func main() {
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(37.2)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
