package main

import "fmt"

func main() {
	arr2 := [5]int{1,3,5,8,7}
	sum := 8
	for i:=0; i<len(arr2); i++ {
//		fmt.Println(value)
		for j:=0; j<len(arr2); j++ {
			//fmt.Println(arr2[i])
			if arr2[j] + arr2[i] == sum {
				fmt.Println(i,j)
			}
		}
	}
}
