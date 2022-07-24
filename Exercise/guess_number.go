//https://www.cnblogs.com/forever521Lee/p/9338157.html
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var input int
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1000)
	//fmt.Println(num)

	for i := 0; i < 10; i++ {
		fmt.Printf("有%d次机会", 10-i)
		fmt.Print("请输入一个数字：")
		fmt.Scan(&input)
		//fmt.Println(input)
		switch {
		case num < input:
			fmt.Println("big")
		case num > input:
			fmt.Println("small")
		case num == input:
			fmt.Println("right")
			return
		default:
			continue
		}
	}
}
