package main

import "fmt"

func (cli *CLI) walletAddress() {
	wallet := GetWallet()
	fmt.Printf("Wallet Bytes Address: %x\n", wallet.GetAddress())
	fmt.Printf("Wallet Readable Address : %s\n", wallet.GetAddress())
}
