package main

import "fmt"

func main() {
	s1 := []string{"s1", "s1"}
	s1 = append(s1, "s2", "s2")
	s1 = append(s1, "s3", "s3")
	s1 = append(s1, "s4", "s4")
	fmt.Println(s1)
}
