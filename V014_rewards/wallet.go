package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "wallet.dat"

// Wallet ...
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
}

func walletIsExits() bool {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// NewWallet creates and returns a Wallet
func NewWallet() *Wallet {
	if walletIsExits() {
		fmt.Println("Wallet already exists")
		os.Exit(1)
	}
	private, public := newKeyPair()
	wallet := Wallet{
		PrivateKey: private,
		PublicKey:  public}

	return &wallet
}

func newKeyPair() (ecdsa.PrivateKey, ecdsa.PublicKey) {
	curve := elliptic.P256()
	//private, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	private, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		log.Panic(err)
	}

	return *private, private.PublicKey
}

// SaveToFile saves the wallet to a file
func (wallet Wallet) SaveToFile() {
	var content bytes.Buffer

	// add this line
	gob.Register(wallet.PrivateKey.Curve)

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(wallet)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}

// GetWallet ...
func GetWallet() *Wallet {
	if !walletIsExits() {
		fmt.Println("Wallet not exists")
		os.Exit(1)
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallet Wallet

	//add this two lines
	curve := elliptic.P256()
	gob.Register(curve)

	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallet)

	if err != nil {
		log.Panic(err)
	}

	return &wallet

}
