package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade := grades[name]
	if grade < 60 {
		fmt.Printf("%s is failling!\n", name)
	}
}

func main() {
	status("Alma")
	status("Carl")
}
