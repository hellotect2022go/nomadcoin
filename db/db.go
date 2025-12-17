package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/hellotect2022go/nomadcoin/utils"
)

const (
	dbName       = "blockChain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			//utils.HandleErr(err)
			fmt.Println("dataBucket 새성시 에러?", err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			fmt.Println("blocksBucket 새성시 에러?", err)
			return err
		})
		utils.HandleErr(err)

	}
	return db
}

func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte("checkpoint"), data)
		return err
	})
	utils.HandleErr(err)
}
