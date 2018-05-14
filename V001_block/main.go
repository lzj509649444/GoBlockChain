package main

import (
	"fmt"
)

func main() {
	genesisBlock := NewGenesisBlock()
	block := NewBlock("Send 1 BTC to Ivan", genesisBlock.Hash)

	fmt.Printf("%s\n", genesisBlock.Data)
	fmt.Printf("%x\n", genesisBlock.Hash)
	fmt.Printf("%s\n", block.Data)
	fmt.Printf("%x\n", block.Hash)
}
