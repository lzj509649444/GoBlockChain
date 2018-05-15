package main

import "bytes"

// TXOutput represents a transaction output
type TXOutput struct {
	Value        int
	ScriptPubKey []byte
	//ScriptPubKey string
}

// Unlock checks if the output can be used by the owner of the pubkey
func (out *TXOutput) Unlock(pubKeyHash []byte) bool {
	return bytes.Compare(out.ScriptPubKey, pubKeyHash) == 0
}

// LockOutput signs the output
func (out *TXOutput) LockOutput(address string) {
	out.ScriptPubKey = GetHashPubKey(address)
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) TXOutput {
	out := TXOutput{
		Value:        value,
		ScriptPubKey: nil}
	out.LockOutput(address)

	return out
}
