package main

import "fmt"

func main() {
	numbers := make(map[string]string)
	numbers["I've been assigned"] = "hi"
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I haven't been assigned"])
}
