package main

import (
	"fmt"
	//"strconv"
	"math"
)

func main() {
	fmt.Println(isPalindrome(2))
}

func isPalindrome(x int) bool {
	if x == 1{
		return true
	}
	num_slice := []int{}
	for i := 1; i < x; i *= 10 {
		num_slice = append(num_slice, x/i % 10)
	}
	var tmp_num float64
	length := len(num_slice)
	for i:= 0; i < length; i++ {
		tmp_num += float64(num_slice[i]) * math.Pow(10,float64(length - i - 1))
	}
	fmt.Println("result: %s",tmp_num)
	if float64(x) == tmp_num {
		return true
		
	}else {
		return false
	}

}
