package main

import "fmt"

func printValue(cityMap map[string]string) {
    for key,value := range cityMap {
        fmt.Println(key, value)
    }
}

func updateValue(cityMap map[string]string) {
    cityMap["fujian"] = "xiamen"
}

func main() {
    cityMap := make(map[string]string)
    // add
    cityMap["guangdong"] = "guangzhou"
    cityMap["fujian"] = "fuzhou"
    cityMap["shanxi"] = "xian"

    //list 
    fmt.Println("--------------------")
    printValue(cityMap)
    
    //update
    fmt.Println("--------------------")
    updateValue(cityMap)
    printValue(cityMap)

    //delete
    fmt.Println("--------------------")
    delete(cityMap, "fujian")
    printValue(cityMap)
}
    
