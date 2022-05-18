package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Yekai")
	bc.AddBlock("Send 2 Btc to Fuhongxue")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %t\n", pow.Validate())
		fmt.Println()
	}
}
