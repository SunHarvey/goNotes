package main

import "fmt"

type Person struct {
	Name        string
	Gender, Age int
}

type Employee struct {
	p      Person
	Salary int
}

type Student struct {
	Person
	School string
}

func main() {
	e := Employee{p: Person{"Soctt", 1, 30}, Salary: 1000}
	fmt.Println(e)
	fmt.Println(e.p.Name)

	var s Student
	s.Name = "Billy"
	s.Gender = 1
	s.Age = 6
	s.School = "梦想小学"
	fmt.Println(s)
}
