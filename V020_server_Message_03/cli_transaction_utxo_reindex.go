package main

import (
	"fmt"
)

func (cli *CLI) reindexUTXO() {
	blockchain := GetBlockchain()

	utxoSet := UTXOSet{blockchain}
	utxoSet.Reindex()

	count := utxoSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
