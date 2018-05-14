package main

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func newBucket(tx *bolt.Tx) []byte {
	genesis := NewGenesisBlock()

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

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		if bucket == nil {
			tip = newBucket(tx)
		} else {
			tip = bucket.Get([]byte("l"))
		}

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

func saveBlock(blockchain *Blockchain, newBlock *Block) {
	err := blockchain.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blocksBucket))

		errr := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if errr != nil {
			return errr
		}

		errr = bucket.Put([]byte("l"), newBlock.Hash)
		if errr != nil {
			return errr
		}

		blockchain.tip = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

// AddBlock saves provided data as a block in the blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	lastHash := getLastHash(blockchain)
	newBlock := NewBlock(data, lastHash)
	saveBlock(blockchain, newBlock)
}
