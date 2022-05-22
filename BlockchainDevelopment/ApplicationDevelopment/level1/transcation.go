package main

const subsidy = 10

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type TXInput struct {
	Txid     []byte
	VoutIdx  int
	FromAddr string
}

type TXOutput struct {
	Value  int
	ToAddr string
}

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	enc.Encode(tx)
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}


func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 &&
		len(tx.Vin[0].Txid) == 0 &&
		tx.Vin[0].VoutIdx == -1
}
