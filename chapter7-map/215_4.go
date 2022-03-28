package main

import "fmt"

func main() {
	//var nilMap map[int]string
	var myMap map[int]string = make(map[int]string)
	myMap[3] = "three"
	fmt.Printf("%#v\n", myMap)
}
