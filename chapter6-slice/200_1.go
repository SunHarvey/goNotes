package main

import "fmt"

func severalInt(numbers ...int) {
	fmt.Println(numbers)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	intSlice := []int{1, 2, 3}
	severalInt(intSlice...)
	stringSlice := []string{"a", "b", "c", "d"}
	mix(1, true, stringSlice...)
}
