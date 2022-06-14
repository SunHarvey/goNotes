package main

import "fmt"

func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	fmt.Println("Grades:")
	for _, grade := range grades {
		fmt.Println(grade)
	}
}
