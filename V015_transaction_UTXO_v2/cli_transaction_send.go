package main

import "fmt"

func (cli *CLI) send(from, to string, amount int) {
	blockchain := GetBlockchain()
	defer blockchain.db.Close()

	tx := NewUTXOTransaction(from, to, amount, blockchain)
	rewardsTx := NewCoinbaseTX(from, "")
	blockchain.MineBlock([]*Transaction{tx, rewardsTx})
	fmt.Println("Success!")
}
