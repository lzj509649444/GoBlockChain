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
if nodeAddress != knownNodes[0] {
  sendVersion(knownNodes[0], blockchain)
}

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


Starting node 4000
Received version command
knownNodes [localhost:4000 localhost:4001]
Received getblocks command
Received getdata command
Received getdata command
Received version command
knownNodes [localhost:4000 localhost:4001 localhost:4002]
Received getblocks command
Received getdata command
Received getdata command

Starting node 4001
Received version command
knownNodes [localhost:4000]
Received inv command
Recevied inventory with 2 blocks
Received block command
Recevied a new block!
[[0 3 200 46 181 255 250 230 81 12 60 117 14 83 14 49 27 240 143 150 50 135 240 5 140 22 227 57 31 238 46 54]]
Received block command
Recevied a new block!
[]

Starting node 4002
Received version command
knownNodes [localhost:4000]
Received inv command
Recevied inventory with 2 blocks
Received block command
Recevied a new block!
[[0 3 200 46 181 255 250 230 81 12 60 117 14 83 14 49 27 240 143 150 50 135 240 5 140 22 227 57 31 238 46 54]]
Received block command
Recevied a new block!
[]
