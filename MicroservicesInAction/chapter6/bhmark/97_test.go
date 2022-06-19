package exam

import(
	"fmt"
	_ "testing"
)

func Sum(a, b int) int {
	return a + b
}

func ExampleSum() {
	fmt.Println(Sum(1,2))
}

