package main

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}
