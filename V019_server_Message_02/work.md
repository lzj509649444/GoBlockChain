https://en.bitcoin.it/wiki/Protocol_documentation#inv

Messages: version,verack,addr

1. block Height,MineBlock

2. Version,sendVersion
BestHeight int

if nodeID != dnsNodeID {
  sendVersion(fmt.Sprintf("localhost:%d", dnsNodeID), blockchain)
}

3. handleVersion, BestHeight
4. server.go
5. server_message_02.go

6. startnode cli

NODE_ID=4000 ./main createwallets

NODE_ID=4000 ./main addwallet
wallet address:  1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

NODE_ID=4000 ./main addwallet
wallet address:  1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez

NODE_ID=4000 ./main addwallet
wallet address:  17pNW3EuirpDDXyrY5oqBCuMvRNL27uKjk


NODE_ID=4000 ./main createblockchain -address 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

NODE_ID=4000 ./main getbalance -address 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

cp blockchain_4000.db blockchain_genesis.db

NODE_ID=4000 ./main send -from 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1 -to 1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez -amount 3

NODE_ID=4000 ./main getbalance -address 1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez

NODE_ID=4000 ./main startnode

# open a new shell window
cp blockchain_genesis.db blockchain_4001.db
NODE_ID=4001 ./main startnode
stopnode and check
NODE_ID=4001 ./main printchain
