package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println("cap:", cap(a))

	//a = append(a[:1], 222)
	a = append(a[:2], append([]int{-222}, a[2:]...)...)
	fmt.Println(a)
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("cap:", cap(a))
}
