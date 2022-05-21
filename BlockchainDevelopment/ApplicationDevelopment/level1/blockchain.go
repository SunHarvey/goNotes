package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	//	blocks []*Block
	tip []byte
	db  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var tip []byte
	bc.db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(blocksBucket))
		tip = buck.Get([]byte("last"))
		return nil
	})

	bc.db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(blocksBucket))
		block := NewBlock(data, tip)

		buck.Put(block.Hash, block.Serialize())
		buck.Put([]byte("last"), block.Hash)
		bc.tip = block.Hash
		return nil
	})
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, _ := bolt.Open(dbFile, 0600, nil)
	db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(blocksBucket))
		if buck == nil {
			fmt.Println("No existing blockchain found. Creating a new one ...")
			genesis := NewGenesisBlock()
			block_data := genesis.Serialize()
			bucket, _ := tx.CreateBucket([]byte(blocksBucket))
			bucket.Put(genesis.Hash, block_data)
			bucket.Put([]byte("last"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = buck.Get([]byte("last"))
		}
		return nil
	})
	return &Blockchain{tip, db}
}
