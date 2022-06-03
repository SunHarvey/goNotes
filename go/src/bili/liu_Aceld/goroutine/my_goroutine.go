package main

import (
    "fmt"
    "time"
)

func hello_world(num int) {
    fmt.Println("hello world", num)
}
func main() {
    for i:=0; i<5; i++ {
        go hello_world(i)
    }
    time.Sleep(1 * time.Second)
}
    
    
