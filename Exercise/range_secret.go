package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	numCharset = "0123456789"
	strCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	mixCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	advCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`~!@#$%^&*()-_+=|,.<>?/"
)

func main() {
	var length int
	var charset string
	flag.IntVar(&length, "l", 16, "-l the length of passwd")
	flag.StringVar(&charset, "t", "harden", "-t the charset of password")
	flag.Parse()
	fmt.Println(length, charset)
	var userCharset string
	switch charset {
	case "num":
		userCharset = numCharset
	case "str":
		userCharset = strCharset
	case "mix":
		userCharset = mixCharset
	case "adv":
		userCharset = advCharset
	default:
		userCharset = mixCharset
	}

	var password []byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(userCharset))
		char := userCharset[index]
		password = append(password, char)
	}

	strPassword := string(password)
	fmt.Printf("%s\n", strPassword)
}
