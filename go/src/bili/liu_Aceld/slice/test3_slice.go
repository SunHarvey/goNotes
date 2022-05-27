package main

import "fmt"

func main() {
	//slice1 := []int{1,2,3}

	var slice1 []int
	slice1 = make([]int, 3)
	slice1[0] = 100
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	var slice2 []int
	if slice2 == nil {
		fmt.Println("empty")
	} else {
		fmt.Println(" no empty")
	}
}
