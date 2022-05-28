package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) SetName() {
	h.Name = "abc"
}

func (h *Hero) Show() {
	fmt.Println(h.Name)
	fmt.Println(h.Ad)
	fmt.Println(h.Level)
}
func main() {
	hero1 := Hero{Name: "liubang", Ad: 100, Level: 2}
	//fmt.Println(hero1.Name)
	fmt.Println("hero1 name : ", hero1.GetName())
	hero1.SetName()
	fmt.Println("hero1 name : ", hero1.GetName())
	hero1.Show()

}
