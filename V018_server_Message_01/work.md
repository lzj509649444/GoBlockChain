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

3. blockchain NodeID name
4. cli: func getNodeID()

NODE_ID=4000 ./main createblockchain -address 1FgRHeYoRXEnJ2YqtYwEkZMcATgDWkv8XL

NODE_ID=4000 ./main getbalance -address 1FgRHeYoRXEnJ2YqtYwEkZMcATgDWkv8XL

NODE_ID=4000 ./main send -from 1FgRHeYoRXEnJ2YqtYwEkZMcATgDWkv8XL -to 1Hu2xHZV97GSTDezZf7mtLfaXXwjPHcV6a -amount 3
