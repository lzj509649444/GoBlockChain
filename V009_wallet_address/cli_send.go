package main

import "fmt"

func (cli *CLI) send(from, to string, amount int) {
	blockchain := GetBlockchain()
	defer blockchain.db.Close()

	tx := NewUTXOTransaction(from, to, amount, blockchain)
	blockchain.MineBlock([]*Transaction{tx})
	fmt.Println("Success!")
}
