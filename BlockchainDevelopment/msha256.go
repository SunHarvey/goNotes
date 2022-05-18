package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	str := os.Args[1]
	sum := sha256.Sum256([]byte(str))
	fmt.Printf("%x\n", sum) //a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
}
