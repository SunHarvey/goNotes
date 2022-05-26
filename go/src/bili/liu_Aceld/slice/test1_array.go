package main

import "fmt"

func main() {
    var myArray1 [10]int
    myArray2 := [10]int{1,2,3,4}
    
    for i:=0;i<10;i++{
        fmt.Println(myArray1[i])
    }

    for index,value := range myArray2 {
        fmt.Println("index = ", index, ",value = ",value)
    }
    
    fmt.Printf("myArray1 type = %T\n", myArray1)
    fmt.Printf("myArray2 type = %T\n", myArray2)
}
