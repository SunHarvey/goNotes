package main

import "fmt"


func main() {
    defer fmt.Println("main end")
    fmt.Println("main:: hello go 1")
    fmt.Println("main:: hello go 2")
}
