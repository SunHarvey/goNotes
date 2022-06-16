package main

import "fmt"

func main() {
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	fmt.Println(u8, u16, u32, u64)

	var p uintptr
	p = 1888888888888889991
	fmt.Printf("p: %v\n", p)
}
