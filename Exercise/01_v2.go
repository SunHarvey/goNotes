package main

import "fmt"

func main() {
    var m, n, Sn, Tn = 1, 1, 0, 0

    fmt.Print("请输入一个数字：")
    fmt.Scan(&m)
    fmt.Print("请输入位数：")
    fmt.Scan(&n)
    for i := 1; i <= n; i++ {
        Tn = Tn * 10
        Tn = Tn + m
        Sn = Sn + Tn

    }
    fmt.Println("这个结果是：", Tn)
}
