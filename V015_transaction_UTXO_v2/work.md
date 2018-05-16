1. TXOutputs

2. UTXOSet
   func Reindex()
	 func Update()
	 func (blockchain *Blockchain) FindUTXO() map[string]TXOutputs
   func (utxoSet UTXOSet) FindSpendableOutputs(pubkeyHash []byte, amount int) (int, map[string][]int)
	 func NewUTXOTransaction(from, to string, amount int, UTXOSet *UTXOSet)

	 func (utxoSet UTXOSet) CountTransactions() int

3. wallet_address
   func ValidateAddress(address string) bool

4. cli_blockchain_create
   UTXOSet := UTXOSet{blockchain}
   UTXOSet.Reindex()

5. cli_transaction_getbalance
   utxoSet := UTXOSet{Blockchain: blockchain}

6. cli_send
   utxoSet := &UTXOSet{Blockchain: blockchain}

   update blockchain.MineBlock
   utxoSet.Update(block)

7. cli_transaction_utxo_reindex


./main addwallet
wallet address:  1FgRHeYoRXEnJ2YqtYwEkZMcATgDWkv8XL
Add Wallet Success !

 ./main addwallet
wallet address:  1F2WH1wD2eGKkoYG74czFmCbE1qBaGHCXP
Add Wallet Success !

./main addwallet
wallet address:  1Hu2xHZV97GSTDezZf7mtLfaXXwjPHcV6a
Add Wallet Success !

./main createblockchain -address 1FgRHeYoRXEnJ2YqtYwEkZMcATgDWkv8XL

./main send -from 1FgRHeYoRXEnJ2YqtYwEkZMcATgDWkv8XL -to 1Hu2xHZV97GSTDezZf7mtLfaXXwjPHcV6a -amount 3

./main send -from 1Hu2xHZV97GSTDezZf7mtLfaXXwjPHcV6a -to 1F2WH1wD2eGKkoYG74czFmCbE1qBaGHCXP -amount 2
