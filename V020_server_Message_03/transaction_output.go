package main

import (
	"bytes"
	"fmt"
	"strings"
)

// TXOutput represents a transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// String returns a human-readable representation of a transaction
func (out TXOutput) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("  Output:"))
	lines = append(lines, fmt.Sprintf("    Value: %d", out.Value))
	lines = append(lines, fmt.Sprintf("    Script: %x} ", out.PubKeyHash))
	return strings.Join(lines, "\n")
}

// Unlock checks if the output can be used by the owner of the pubkey
func (out *TXOutput) Unlock(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// LockOutput signs the output
func (out *TXOutput) LockOutput(address string) {
	out.PubKeyHash = GetHashPubKey(address)
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) TXOutput {
	out := TXOutput{
		Value:      value,
		PubKeyHash: nil}
	out.LockOutput(address)

	return out
}
