package main

import (
	"fmt"
	"github.com/headfirstgo/prose"
)

func main() {
	phrases := []string{"my parents","a rodo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
