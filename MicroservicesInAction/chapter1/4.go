package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := "你好，世界！"
	fmt.Fprintf(w, "%s", s)
	log.Printf("%s", s)
}

func main() {
	fmt.Println("server start.")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:1234", nil)
	if err != nil {
		log.Fatal("ListenAdnServer:", err)
	}
}
