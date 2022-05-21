package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func main() {
	db, _ := bolt.Open("my.db", 0600, nil)
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bucket,_ := tx.CreateBucket([]byte("bucket1"))
		bucket.Put([]byte("name"), []byte("yekai"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("bucket1"))
		val := bucket.Get([]byte("name"))
		fmt.Println(string(val))
		return nil
	})
}
