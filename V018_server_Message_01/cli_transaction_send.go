package main

import (
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	blockchain := GetBlockchain(getNodeID())
	defer blockchain.db.Close()

	utxoSet := UTXOSet{Blockchain: blockchain}

	tx := NewUTXOTransaction(from, to, amount, &utxoSet)
	rewardsTx := NewCoinbaseTX(from, "")
	block := blockchain.MineBlock([]*Transaction{tx, rewardsTx})
	utxoSet.Update(block)
	fmt.Println("Success!")
}
