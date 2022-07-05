//1.题目：求s=a+aa+aaa+aaaa+aa...a的值，其中a是一个数字。例如2+22+222+2222+22222(此时共有5个数相加)，几个数相加有键盘控制。
//
//2.程序分析：关键是计算出每一项的值。

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var number string
	fmt.Print("input number:")
	fmt.Scanln(&number)
	//fmt.Println(number)
	var num_list []string
	num, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(num)
	}
	for i := 1; i <= num; i++ {
		str_num := strings.Repeat(number, i)
		num_list = append(num_list, str_num)
	}
	fmt.Println(num_list)
	var sum int64 = 0
	for _, value := range num_list {
		//fmt.Printf("%T\n",value)
		n, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			os.Exit(1)
		}
		//fmt.Printf("n type: %T\n",n)
		sum += n

	}
	fmt.Println(sum)

}
