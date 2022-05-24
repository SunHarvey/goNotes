package main

import (
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"log"
)

func create_mnemonic() {
	b, err := bip39.NewEntropy(128)
	if err != nil {
		log.Panic("Failed to NewEntropy:", err, b)
	}
	nm, err := bip39.NewMnemonic(b)
	if err != nil {
		log.Panic("Failed to NewMnemonic:", err)
	}
	fmt.Println(nm)
}

func main() {
	create_mnemonic()
}
