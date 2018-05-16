package main

import (
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	blockchain := GetBlockchain()
	defer blockchain.db.Close()

	utxoSet := UTXOSet{Blockchain: blockchain}

	balance := 0
	UTXOs := utxoSet.FindUTXO(GetHashPubKey(address))

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
