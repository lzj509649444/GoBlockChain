package main

import (
	"fmt"
)

func main() {

	blockChain := NewBlockchain()

	blockChain.AddBlock("Send 1 BTC to Ivan")
	blockChain.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range blockChain.blocks {
		fmt.Printf("%s\n", block.Data)
		fmt.Printf("%x\n", block.hash)
	}
}
