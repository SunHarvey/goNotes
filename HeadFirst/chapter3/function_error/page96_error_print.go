package main

import (
	//	"errors"
	"fmt"
	//	"log"
)

func main() {
	//err := errors.New("height can't be nagative")
	//	fmt.Println(err.Error())
	//	fmt.Println(err)
	//	log.Fatal(err)
	err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	fmt.Println(err.Error())
	fmt.Println(err)
}
