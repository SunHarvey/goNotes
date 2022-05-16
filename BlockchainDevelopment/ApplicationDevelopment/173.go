package main

import (
	"fmt"
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+i, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func Base58Encode(input int64) []byte {
	var result []byte
	x := big.NewInt(input)
	base := big.NewInt(int64(len(b58Alphabet)))

	zero := big.NewInt(0)
	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)
	return result
}

func main() {
	result := Base58Encode(258)
	fmt.Println(string(result))
}
