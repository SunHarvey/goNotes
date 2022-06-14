package main

import (
	"fmt"
	"github.com/headfirstgo/magazine"
)
/*
func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress.Street = "123 Oak St"
	subscriber.HomeAddress.City = "Omaha"
	subscriber.HomeAddress.State = "NE"
	subscriber.HomeAddress.PostalCode = "68111"
	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.HomeAddress.Street)
	fmt.Println("City:", subscriber.HomeAddress.City)
	fmt.Println("State:", subscriber.HomeAddress.State)
	fmt.Println("PostalCode:", subscriber.HomeAddress.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	employee.HomeAddress.Street = "456 Elm St"
	employee.HomeAddress.City = "Portland"
	employee.HomeAddress.State = "OR"
	employee.HomeAddress.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.HomeAddress.Street)
	fmt.Println("City:", employee.HomeAddress.City)
	fmt.Println("State:", employee.HomeAddress.State)
	fmt.Println("PostalCOde:", employee.HomeAddress.PostalCode)
}
*/
func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("PostalCode:", subscriber.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Street = "456 Elm St"
	employee.City ="Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("PostalCode:", employee.PostalCode)
}


