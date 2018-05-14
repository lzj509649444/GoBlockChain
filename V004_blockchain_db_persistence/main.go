package main

import (
	"fmt"
	"strconv"
)

func main() {

	blockChain := NewBlockchain()

	blockChain.AddBlock("Send 1 BTC to Ivan")
	blockChain.AddBlock("Send 2 more BTC to Ivan")

	iter := blockChain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
