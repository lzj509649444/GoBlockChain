package main

import "bytes"

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

// UnlocksOutputWith checks whether the address initiated the transaction
func (in *TXInput) UnlocksOutputWith(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// NewCoinBaseTXInput ...
func NewCoinBaseTXInput(data string) TXInput {
	input := TXInput{
		Txid:      []byte{},
		Vout:      -1,
		Signature: nil,
		PubKey:    []byte(data)}
	return input
}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

// NewTXInput create a new NewTXInput
// ScriptSig存的是公钥
func NewTXInput(txID []byte, outIndex int, wallet Wallet) TXInput {
	input := TXInput{
		Txid:      txID,
		Vout:      outIndex,
		Signature: nil,
		PubKey:    wallet.PublicKeyBytes()}
	return input
}
