package main

import (
	"fmt"
)

func (cli *CLI) createWallet() {
	wallet := NewWallet()
	wallet.SaveToFile()
	fmt.Println("Wallet Create Success !")
}
