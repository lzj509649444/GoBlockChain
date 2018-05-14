package main

import (
	"fmt"
)

func main() {
	genesisBlock := NewGenesisBlock()
	block := NewBlock("Send 1 BTC to Ivan", genesisBlock.hash)

	fmt.Printf("%s\n", genesisBlock.Data)
	fmt.Printf("%x\n", genesisBlock.hash)
	fmt.Printf("%s\n", block.Data)
	fmt.Printf("%x\n", block.hash)
}
