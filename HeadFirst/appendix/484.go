package main

import (
//	"bufio"
	"log"
	"os"
)


func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.OpenFile("aardvark.txt", os.O_WRONLY, os.FileMode(0600))
	check(err)
	_,err = file.Write([]byte("amazing!\n"))
	check(err)
	err = file.Close()
	check(err)
}


