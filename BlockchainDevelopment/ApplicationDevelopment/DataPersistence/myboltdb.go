package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func myPut(filename string, bname string, key string, value string) {
	db, _ := bolt.Open(filename, 0600, nil)
	fmt.Println(bname, key, value)
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bucket, _ := tx.CreateBucket([]byte(bname))
		bucket.Put([]byte(key), []byte(value))
		return nil
	})
}

func myView(filename string, bname string, key string) {
	db, _ := bolt.Open(filename, 0600, nil)
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bname))
		value := bucket.Get([]byte(key))
		fmt.Println(string(value))
		return nil
	})
}

func main() {
	file := "db.db"
	bname := string("bname3")
	key := "aaaaaa"
	value := "bbbbbbbbb"
	myPut(file, bname, key, value)
	myView(file, bname, key)
}
