package main

import (
	"log"

	"github.com/boltdb/bolt"
)

// BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Iterator ...
func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	iter := &BlockchainIterator{
		currentHash: blockchain.tip,
		db:          blockchain.db}

	return iter
}

// Next returns next block starting from the tip
func (iter *BlockchainIterator) Next() *Block {
	var block *Block

	err := iter.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bucket.Get(iter.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	iter.currentHash = block.PrevBlockHash

	return block
}
