package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
)

// TXOutputs collects TXOutput
type TXOutputs struct {
	Outputs []TXOutput
}

// String returns a human-readable representation of a transaction
func (txoutputs TXOutputs) String() string {
	var lines []string

	for i, output := range txoutputs.Outputs {
		lines = append(lines, fmt.Sprintf("  Output %d:", i))
		lines = append(lines, fmt.Sprintf("    Value: %d", output.Value))
		lines = append(lines, fmt.Sprintf("    Script: %x} ", output.PubKeyHash))
	}

	return strings.Join(lines, "\n")
}

// Serialize serializes TXOutputs
func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// DeserializeOutputs deserializes TXOutputs
func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
