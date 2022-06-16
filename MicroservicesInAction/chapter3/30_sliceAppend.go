package main

import "fmt"

func main() {
    var a = []int{1,2,3}
    fmt.Println("cap:",cap(a))
    a = append(a, 333)
    fmt.Println("cap:",cap(a))
    a = append(a, []int{-333,-333,-333,-333}...)
    fmt.Println("cap:",cap(a))

    for i, v := range a {
        fmt.Println(i,v)
    }
    
    a = append(a[:0],a[:3]...)
    for i,v := range a{
        fmt.Println(i,v)
    }
    fmt.Println("cap:", cap(a))
    a = append([]int{222,222},a...)
    a = append(a[:2], append([]int{-222},a[2:]...)...)
    for i,v := range a{
        fmt.Println(i,v)
    }
    fmt.Println("cap:", cap(a))
}

