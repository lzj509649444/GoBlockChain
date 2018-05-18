package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
	"time"
)

const subsidy = 10

// Transaction represents a Bitcoin transaction
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// String returns a human-readable representation of a transaction
func (transaction Transaction) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("{ID: %x", transaction.ID))

	for i, input := range transaction.Vin {

		lines = append(lines, fmt.Sprintf("  Input %d:", i))
		lines = append(lines, fmt.Sprintf("    TXID: %x", input.Txid))
		lines = append(lines, fmt.Sprintf("    Out: %d", input.Vout))
		lines = append(lines, fmt.Sprintf("    Signature: %x", input.Signature))
		lines = append(lines, fmt.Sprintf("    PubKey: %x", input.PubKey))
	}

	for i, output := range transaction.Vout {
		lines = append(lines, fmt.Sprintf("  Output %d:", i))
		lines = append(lines, fmt.Sprintf("    Value: %d", output.Value))
		lines = append(lines, fmt.Sprintf("    Script: %x} ", output.PubKeyHash))
	}

	return strings.Join(lines, "\n")
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	// if data == "" {
	// 	data = fmt.Sprintf("Reward to '%s'", to)
	// }

	if data == "" {
		data = fmt.Sprintf("%x", time.Now())
		fmt.Println("data: ", data)
	}

	txinput := NewCoinBaseTXInput(data)
	txoutput := NewTXOutput(subsidy, to)

	tx := Transaction{
		ID:   nil,
		Vin:  []TXInput{txinput},
		Vout: []TXOutput{txoutput}}
	//tx.SetID()
	tx.ID = tx.Hash()

	return &tx
}

// DeserializeTransaction deserializes a transaction
func DeserializeTransaction(data []byte) Transaction {
	var transaction Transaction

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&transaction)
	if err != nil {
		log.Panic(err)
	}

	return transaction
}
