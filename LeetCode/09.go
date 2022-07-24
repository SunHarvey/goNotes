package main

import (
	"fmt"
	"strconv"
)

func main() {
	//str1 := strconv.Itoa(1000)
	//fmt.Printf("type: %T, value: %#v\n", str1,str1)
	fmt.Println(12321 % 10)
	fmt.Println(12321 % 100)
	fmt.Println(12321 % 1000)

	fmt.Println(isPalindrome(12321))
}

func isPalindrome(x int) bool {
	tmp_x := strconv.Itoa(x)
	fmt.Printf("type: %T, value: %#v\n", tmp_x, tmp_x)
	length := len(tmp_x)
	for i := 0; i < length; i++ {
		fmt.Println(string(tmp_x[i]))
		if tmp_x[i] == tmp_x[length-i-1] {
			return false
		}
	}
	return true
}
