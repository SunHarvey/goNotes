package main

import "fmt"

func calmDown() {
	p := recover()
	fmt.Println(p.Error())
}

func main() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err)
}
