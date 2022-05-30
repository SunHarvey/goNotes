package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Tittle string   `jaon:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"xijuzhiwang", 2000, 10, []string{"xingye", "zhangbozhi"}}

	// encode, struct --> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// decode jsonstr --> struct
	// jsonStr = {"Tittle":"xijuzhiwang","year":2000,"rmb":10,"actors":["xingye","zhangbozhi"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("Json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
}
