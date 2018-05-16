package main

import "fmt"

func (cli *CLI) walletAddress() {
	wallet := GetWallet()
	fmt.Printf("Wallet Address: %x\n", wallet.GetAddress())
}
