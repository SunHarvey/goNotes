package main

import (
	"fmt"
	"github.com/headfirstgo/magazine"
)

func main() {
	subscriber := magazine.Subscriber{}
	fmt.Printf("%#v\n", subscriber.HomeAddress)

	subscriber.HomeAddress.City = "New York"
	fmt.Println(subscriber.HomeAddress.City)

	subscriber.HomeAddress.PostalCode = "68111"
	fmt.Println(subscriber.HomeAddress.PostalCode)
}
