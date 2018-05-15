package main

import "bytes"

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig []byte
	//ScriptSig string
}

// CanUnlockOutputWith checks whether the address initiated the transaction
// func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
// 	return in.ScriptSig == unlockingData
// }

// UnlocksOutputWith checks whether the address initiated the transaction
func (in *TXInput) UnlocksOutputWith(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.ScriptSig)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// NewTXInput create a new NewTXInput
func NewTXInput(txID []byte, outIndex int, wallet Wallet) TXInput {
	input := TXInput{
		Txid:      txID,
		Vout:      outIndex,
		ScriptSig: wallet.PublicKeyBytes()}
	return input
}
