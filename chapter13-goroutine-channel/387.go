package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	var size int
	size = go responseSize("https://baidu.com")
	fmt.Println(size)
	size = go responseSize("https://go.dev")
	fmt.Println(size)
	size = go responseSize("https://go.dev/doc")
	fmt.Println(size)

	time.Sleep(4 * time.Second)
}

func responseSize(url string) int {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(len(body))
	return len(body)
}
