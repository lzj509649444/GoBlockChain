package main

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "genesis coin base"

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

func newBucket(address string, tx *bolt.Tx) []byte {

	coinBaseTx := NewCoinbaseTX(address, genesisCoinbaseData)
	genesis := NewGenesisBlock(coinBaseTx)

	newBucket, err := tx.CreateBucket([]byte(blocksBucket))
	if err != nil {
		log.Panic(err)
	}

	err = newBucket.Put(genesis.Hash, genesis.Serialize())
	if err != nil {
		log.Panic(err)
	}

	err = newBucket.Put([]byte("l"), genesis.Hash)
	if err != nil {
		log.Panic(err)
	}
	return genesis.Hash
}

// CreateBlockchain creates a new blockchain DB
func CreateBlockchain(address string) *Blockchain {
	if dbExists() {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		tip = newBucket(address, tx)
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	blockchain := Blockchain{tip, db}

	return &blockchain

}

// GetBlockchain creates a new Blockchain with genesis Block
func GetBlockchain() *Blockchain {

	if dbExists() == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		tip = bucket.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	blockchain := Blockchain{tip, db}

	return &blockchain
}

func getLastHash(blockchain *Blockchain) []byte {
	var lastHash []byte

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte("l"))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	return lastHash
}

// MineBlock mines a new block with the provided transactions
func (blockchain *Blockchain) MineBlock(transactions []*Transaction) *Block {

	for _, tx := range transactions {
		if blockchain.VerifyTransaction(tx) != true {
			log.Panic("ERROR: Invalid transaction")
		}
	}

	lastHash := getLastHash(blockchain)
	newBlock := NewBlock(transactions, lastHash)

	err := blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		errr := b.Put(newBlock.Hash, newBlock.Serialize())
		if errr != nil {
			return errr
		}

		errr = b.Put([]byte("l"), newBlock.Hash)
		if errr != nil {
			return errr
		}

		blockchain.tip = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	return newBlock
}
