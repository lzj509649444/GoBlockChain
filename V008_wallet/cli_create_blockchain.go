package main

import "fmt"

func (cli *CLI) createBlockchain(address string) {
	blockchain := CreateBlockchain(address)
	blockchain.db.Close()
	fmt.Println("Done!")
}
