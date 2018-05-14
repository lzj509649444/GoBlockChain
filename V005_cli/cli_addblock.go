package main

import "fmt"

func (cli *CLI) addBlock(blockchain *Blockchain, data string) {
	blockchain.AddBlock(data)
	fmt.Println("Success!")
}
