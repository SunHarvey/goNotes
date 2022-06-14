package main

import (
	"fmt"
	"log"
)

type OverheartError float64

func (o OverheartError) Error() string {
	return fmt.Sprintf("Overhearting by %0.2f degress!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheartError(excess)
	}
	return nil
}

func main() {
	var err error = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
