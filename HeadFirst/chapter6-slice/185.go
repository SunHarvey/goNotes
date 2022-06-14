package main

import "fmt"

func main() {
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)

	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
}
