package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func responseSize(url string, channel chan string) {
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
	//channel <- string(len(body)) + url
	channel <- strconv.Itoa(len(body)) + ": " + url
}

func main() {
	sizes := make(chan string)
	urls := []string{"https://www.baidu.com", "https://www.qq.com", "https://www.jianshu.com"}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
}
