https://en.bitcoin.it/wiki/Protocol_documentation#inv

Messages: version,verack,addr

1. sendTx
2. handleTx
3. handleInv: type == 'tx'
4. handleGetData: type == 'tx'
5. send: -mine
cli.send(*sendFrom, *sendTo, *sendAmount, *sendMine)
6. server_message_02.go: handleBlock
utxoSet := UTXOSet{blockchain}
utxoSet.Reindex()

7. cli startnode: -miner
StartServer(nodeID, minerAddress)
