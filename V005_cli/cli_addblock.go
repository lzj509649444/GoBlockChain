package main

import "fmt"

func (cli *CLI) addBlock(data string) {
	blockchain := NewBlockchain()
	blockchain.AddBlock(data)
	fmt.Println("Success!")
}
