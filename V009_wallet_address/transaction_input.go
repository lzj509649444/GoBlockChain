package main

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// CanUnlockOutputWith checks whether the address initiated the transaction
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
