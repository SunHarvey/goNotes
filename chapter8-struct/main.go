package main

import (
	"fmt"
	"github.com/headfirstgo/magazine"
)

func main() {
	/*
			var s magazine.Subscriber
			s.Rate = 4.99
			fmt.Println(s.Rate)
		var employee magazine.Employee
		employee.Name = "Joy Carr"
		employee.Salary = 6000
		fmt.Println(employee.Name)
		fmt.Println(employee.Salary)
	*/
	var address magazine.Address
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"
	fmt.Println(address)
}
