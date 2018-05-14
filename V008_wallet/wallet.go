package main

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const version = byte(0x00)
const walletFile = "wallet.dat"

// Wallet ...
type Wallet struct {
	PrivateKey []byte
	PublicKey  []byte
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

func newKeyPair() ([]byte, []byte) {
	curve := elliptic.P256()

	private, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	public := append(x.Bytes(), y.Bytes()...)

	return private, public
}

// SaveToFile saves the wallet to a file
func (wallet Wallet) SaveToFile() {
	var content bytes.Buffer

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
