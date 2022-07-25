//https://www.cnblogs.com/forever521Lee/p/9338157.html#_label1
package main

import "fmt"

func main() {
	list_a :=[8]int{1,3,5,7,9,11,13,15}
	sum := 0
	for _,value := range list_a {
		//fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)
}
