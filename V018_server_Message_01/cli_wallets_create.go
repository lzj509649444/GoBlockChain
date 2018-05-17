package main

import (
	"fmt"
)

func (cli *CLI) createWallets() {
	wallets := NewWallets()
	wallets.SaveToFile()
	fmt.Println("Wallets Create Success !")
}
