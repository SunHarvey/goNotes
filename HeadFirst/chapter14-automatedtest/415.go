package main

import (
	"fmt"
	"github.com/headfirstgo/prose"
)

func main() {
	phrases := []string{"my parents"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
