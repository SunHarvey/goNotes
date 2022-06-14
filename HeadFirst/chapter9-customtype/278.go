package main

import "fmt"

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Substract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func main() {
	ten := Number(10)
	ten.Add(4)
	ten.Substract(5)
	four := Number(4)
	four.Add(3)
	four.Substract(2)

}
