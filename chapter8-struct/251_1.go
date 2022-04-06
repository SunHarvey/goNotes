package main

import (
	"fmt"
	"github.com/headfirstgo/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)
}
