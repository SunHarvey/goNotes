package main

import "fmt"

func main() {
	counters := map[string]int{"a": 3, "b": 0}
	var ok bool
	_, ok = counters["b"]
	fmt.Println(ok)
	_, ok = counters["c"]
	fmt.Println(ok)
}
