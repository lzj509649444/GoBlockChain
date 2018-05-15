package main

import "fmt"

func (cli *CLI) getBalance(address string) {
	blockchain := GetBlockchain()
	defer blockchain.db.Close()

	balance := 0
	UTXOs := blockchain.FindUTXO(GetHashPubKey(address))

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
