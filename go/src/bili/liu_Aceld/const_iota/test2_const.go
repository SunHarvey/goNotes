package main

import "fmt"

const (
    beijing = iota
    shanghai
    shenzhen
)

func main() {
    const length int = 10
    fmt.Println("length = ", length) 
    fmt.Println(beijing, shanghai,shenzhen)
}




