Messages: version,verack,addr

server_message_01.go

1. start dnsNodeID=4000
./main startnode -node-id 4000

2. ./main startnode -node-id 4001
- nodeAddress=localhost:4001
- sendVersion to dnsNodeID
- dnsNodeID received version command: handleVersion, sendData to ("version{nodeVersion, nodeAddress}"), knownNodes = append(knownNodes, payload.AddrFrom=nodeAddress)
- nodeAddress: Received verack command,Received addr command(handleAddr)
- nodeAddress: append(knownNodes, payload.AddrList...)


4000:
Starting node 4000
Received version command
Received version command

4001:
Starting node 4001
Received verack command
Received addr command
There are 1 known nodes now!

./main startnode -node-id 4002
Starting node 4002
Received verack command
Received addr command
There are 2 known nodes now!

3. cli: func getNodeID()
4. blockchain NodeID name: getDBFile
5. getWalletFile
6. getWalletsFile

NODE_ID=4000 ./main createwallets

NODE_ID=4000 ./main addwallet
wallet address:  1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

NODE_ID=4000 ./main addwallet
wallet address:  1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez

NODE_ID=4000 ./main addwallet
wallet address:  17pNW3EuirpDDXyrY5oqBCuMvRNL27uKjk


NODE_ID=4000 ./main createblockchain -address 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

NODE_ID=4000 ./main getbalance -address 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1

NODE_ID=4000 ./main send -from 1EVT8d6YBvfeE4SYCCRfoWL6ZXk9A3hPU1 -to 1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez -amount 3

NODE_ID=4000 ./main getbalance -address 1Ndq99Y81UuVJ11jwMrMNN17kXMMMYh8Ez
