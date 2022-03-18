package main

import "fmt"

func status(grade float64) string {
    return grade  * 2
}

func main() {
	fmt.Println(status(60.1))
	fmt.Println(status(59.9))
}
