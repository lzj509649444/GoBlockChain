package main

import "fmt"

func (cli *CLI) printWallets() {
	wallets := GetWallets()
	for address, wallet := range wallets.Wallets {
		fmt.Printf("Address: %s\n", address)
		fmt.Printf("Wallet: %+v\n\n", wallet)
	}
}
