1. merkle_tree.go
2. block.go
func (block *Block) HashTransactions() []byte


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
