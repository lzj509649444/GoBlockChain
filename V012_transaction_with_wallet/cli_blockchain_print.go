package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain() {

	blockchain := GetBlockchain()

	iter := blockchain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Block: %+v\n", block)
		block.printTransactions()
		fmt.Printf("PrevBlock. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
