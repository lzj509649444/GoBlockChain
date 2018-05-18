package main

import (
	"fmt"
)

func (cli *CLI) printChain() {

	blockchain := GetBlockchain()

	iter := blockchain.Iterator()

	for {
		fmt.Println("for...")
		block := iter.Next()

		fmt.Printf("Block: %s\n", block)
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
