package main

import (
	"encoding/hex"
	"log"
)

// FindUnspentTransactions returns a list of transactions containing unspent outputs
func (blockchain *Blockchain) FindUnspentTransactions(pubKeyHash []byte) []Transaction {
	var unspentTransactions []Transaction
	spentTransactionOutputs := make(map[string][]int)
	iter := blockchain.Iterator()

	for {
		block := iter.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTransactionOutputs[txID] != nil {
					for _, spentOut := range spentTransactionOutputs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.Unlock(pubKeyHash) {
					unspentTransactions = append(unspentTransactions, *tx)
				}
			}

			// tx.Vin 已消费的交易
			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					// 自己的交易
					if in.UnlocksOutputWith(pubKeyHash) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTransactionOutputs[inTxID] = append(spentTransactionOutputs[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTransactions
}

// FindUTXO finds and returns all unspent transaction outputs
func (blockchain *Blockchain) FindUTXO(pubKeyHash []byte) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := blockchain.FindUnspentTransactions(pubKeyHash)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.Unlock(pubKeyHash) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// FindSpendableOutputs finds and returns unspent outputs to reference in inputs
func (blockchain *Blockchain) FindSpendableOutputs(pubKeyHash []byte, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTransactions := blockchain.FindUnspentTransactions(pubKeyHash)
	accumulated := 0

Work:
	for _, tx := range unspentTransactions {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.Unlock(pubKeyHash) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}

// NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(from, to string, amount int, blockchain *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	wallets := GetWallets()
	wallet := wallets.GetWallet(from)
	pubKeyHash := HashPubKey(wallet.PublicKeyBytes())

	acc, validOutputs := blockchain.FindSpendableOutputs(pubKeyHash, amount)

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
			// input := TXInput{
			// 	Txid:      txID,
			// 	Vout:      out,
			// 	ScriptSig: wallet.PublicKeyBytes()}
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
