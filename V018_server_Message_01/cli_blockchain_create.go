package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	blockchain := CreateBlockchain(address)
	defer blockchain.db.Close()

	utxoSet := UTXOSet{blockchain}
	utxoSet.Reindex()

	fmt.Println("Done!")
}
