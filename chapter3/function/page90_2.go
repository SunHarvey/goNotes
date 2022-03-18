package main

import "fmt"

var metersPerLiter float64

func paintNeeded(width, height float64) float64 {
    area := width * height
    return area / metersPerLiter
}

func main() {
    metersPerLiter = 10.0
    fmt.Printf("%0.2f", paintNeeded(4.2, 3.0))
}
