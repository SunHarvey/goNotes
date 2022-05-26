package main

import "fmt"

func main() {
    cityMap := make(map[string]string)
    // add
    cityMap["guangdong"] = "guangzhou"
    cityMap["fujian"] = "fuzhou"
    cityMap["shanxi"] = "xian"
    
    //list 
    for key,value := range cityMap {
        fmt.Println(key, value)
    }
    fmt.Println("----------------------")
    
    //update
    cityMap["fujian"] = "xiamen"
    for key,value := range cityMap {
        fmt.Println(key, value)
    }

    fmt.Println("----------------------")
    //delete
    delete(cityMap, "fujian")
    for key,value := range cityMap {
        fmt.Println(key, value)
    }
}
    
