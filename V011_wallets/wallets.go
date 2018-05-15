package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletsFile = "wallets.dat"

// Wallets stores a collection of wallets
type Wallets struct {
	Wallets map[string]*Wallet
}

// NewWallet adds a Wallet to Wallets
func (wallets *Wallets) NewWallet() string {
	wallet := NewWallet()
	// 钱包字节地址格式化为string
	address := fmt.Sprintf("%s", wallet.GetAddress())
	wallets.Wallets[address] = wallet
	return address
}

// NewWallets ...
func NewWallets() *Wallets {
	if walletsIsExits() {
		fmt.Println("Wallets already exists")
		os.Exit(1)
	}
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	return &wallets
}

func walletsIsExits() bool {
	if _, err := os.Stat(walletsFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// SaveToFile saves wallets to a file
func (wallets Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(wallets)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletsFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}

// GetWallets loads wallets from the file
func GetWallets() *Wallets {
	if !walletsIsExits() {
		fmt.Println("Wallets not exists !")
		os.Exit(1)
	}

	fileContent, err := ioutil.ReadFile(walletsFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	return &wallets
}
