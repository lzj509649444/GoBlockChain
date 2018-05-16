1. VerifyTransaction
if tx.IsCoinbase() {
	return true
}

2. cli_transaction_send
rewardsTx := NewCoinbaseTX(from, "")
blockchain.MineBlock([]*Transaction{tx, rewardsTx})
