package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (h *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human
	level int
}

func (s *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

func (s *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (s *SuperMan) Print() {
	fmt.Println("name: ",s.name)
	fmt.Println("sex: ",s.sex)
	fmt.Println("level: ", s.level)
}

func main() {
	h := Human{"lisi", "female"}
	h.Eat()
	h.Walk()

	fmt.Println("-------------------------")
	// 1
	//s := SuperMan{Human{"lisi", "female"}, 99}
	// 2
	var s SuperMan
	s.name = "lisi"
	s.sex = "mal"
	s.level = 88

	s.Eat()
	s.Walk()
	s.Fly()
	
	s.Print()
}
