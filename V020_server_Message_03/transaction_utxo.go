package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const utxoBucket = "chainstate"

// UTXOSet represents UTXO set
type UTXOSet struct {
	Blockchain *Blockchain
}

// Reindex rebuilds the UTXO set
func (utxoSet UTXOSet) Reindex() {
	db := utxoSet.Blockchain.db

	bucketName := []byte(utxoBucket)

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucketName)
		if err != nil && err != bolt.ErrBucketNotFound {
			log.Panic(err)
		}

		_, err = tx.CreateBucket(bucketName)
		if err != nil {
			log.Panic(err)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	UTXO := utxoSet.Blockchain.FindUTXO()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)

		for txID, outs := range UTXO {
			key, errr := hex.DecodeString(txID)
			if errr != nil {
				log.Panic(errr)
			}

			fmt.Printf("transaction ID: %x\n", key)
			fmt.Printf("%s\n", outs)
			fmt.Println()

			errr = bucket.Put(key, outs.Serialize())
			if errr != nil {
				log.Panic(errr)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// FindUTXO finds all unspent transaction outputs and returns transactions with spent outputs removed
func (blockchain *Blockchain) FindUTXO() map[string]TXOutputs {
	UTXO := make(map[string]TXOutputs)
	spentTXOs := make(map[string][]int)
	iter := blockchain.Iterator()

	for {
		block := iter.Next()

		for _, tx := range block.Transactions {
			fmt.Printf("transaction ID: %x\n", tx.ID)
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOutIdx := range spentTXOs[txID] {
						if spentOutIdx == outIdx {
							continue Outputs
						}
					}
				}

				fmt.Printf("unspent TXOutput: %s\n", out)

				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return UTXO
}

// FindUTXO finds UTXO for a public key hash
func (utxoSet UTXOSet) FindUTXO(pubKeyHash []byte) []TXOutput {
	var UTXOs []TXOutput
	db := utxoSet.Blockchain.db

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(utxoBucket))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			outs := DeserializeOutputs(v)

			for _, out := range outs.Outputs {
				if out.Unlock(pubKeyHash) {
					UTXOs = append(UTXOs, out)
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return UTXOs
}

// Update updates the UTXO set with transactions from the Block
// The Block is considered to be the tip of a blockchain
func (utxoSet UTXOSet) Update(block *Block) {
	db := utxoSet.Blockchain.db

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(utxoBucket))

		for _, tx := range block.Transactions {
			if tx.IsCoinbase() == false {
				for _, vin := range tx.Vin {
					updatedOuts := TXOutputs{}
					outsBytes := bucket.Get(vin.Txid)
					outs := DeserializeOutputs(outsBytes)

					for outIdx, out := range outs.Outputs {
						if outIdx != vin.Vout {
							updatedOuts.Outputs = append(updatedOuts.Outputs, out)
						}
					}

					if len(updatedOuts.Outputs) == 0 {
						err := bucket.Delete(vin.Txid)
						if err != nil {
							log.Panic(err)
						}
					} else {
						err := bucket.Put(vin.Txid, updatedOuts.Serialize())
						if err != nil {
							log.Panic(err)
						}
					}

				}
			}

			newOutputs := TXOutputs{}
			for _, out := range tx.Vout {
				newOutputs.Outputs = append(newOutputs.Outputs, out)
			}

			err := bucket.Put(tx.ID, newOutputs.Serialize())
			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// FindSpendableOutputs finds and returns unspent outputs to reference in inputs
func (utxoSet UTXOSet) FindSpendableOutputs(pubkeyHash []byte, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	accumulated := 0
	db := utxoSet.Blockchain.db

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(utxoBucket))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			txID := hex.EncodeToString(k)
			outs := DeserializeOutputs(v)

			for outIdx, out := range outs.Outputs {
				if out.Unlock(pubkeyHash) && accumulated < amount {
					accumulated += out.Value
					unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return accumulated, unspentOutputs
}

// CountTransactions returns the number of transactions in the UTXO set
func (utxoSet UTXOSet) CountTransactions() int {
	db := utxoSet.Blockchain.db
	counter := 0

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(utxoBucket))
		cursor := bucket.Cursor()

		for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
			counter++
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return counter
}

// NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(from, to string, amount int, utxoSet *UTXOSet) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	blockchain := utxoSet.Blockchain

	wallets := GetWallets()
	wallet := wallets.GetWallet(from)
	pubKeyHash := HashPubKey(wallet.PublicKeyBytes())

	acc, validOutputs := utxoSet.FindSpendableOutputs(pubKeyHash, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := NewTXInput(txID, out, wallet)
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, NewTXOutput(amount, to))
	if acc > amount {
		outputs = append(outputs, NewTXOutput(acc-amount, from)) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	//tx.SetID()
	tx.ID = tx.Hash()
	blockchain.SignTransaction(&tx, wallet.PrivateKey)

	return &tx
}
