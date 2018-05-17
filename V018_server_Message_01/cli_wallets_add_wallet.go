package main

import (
	"fmt"
)

func (cli *CLI) addWallet() {
	wallets := GetWallets()
	address := wallets.NewWallet()
	fmt.Println("wallet address: ", address)
	wallets.SaveToFile()
	fmt.Println("Add Wallet Success !")
}
