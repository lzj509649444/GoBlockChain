https://en.bitcoin.it/wiki/Protocol_documentation#inv

场景：一个中心节点4000，钱包节点4001，矿工节点4002

从4000和4001 send出去的交易，被4002矿工节点挖出来，startnode各自同步区块

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

8. handleInv: blocks -> block
9. sendData: updatedNodes
10. Sign
11. Verify
12. fix AddBlcok: blockchain.tip = block.Hash
13. Base58
14. fix:创世区块，随机data，避免重复

# 4000 shell
NODE_ID=4000 ./main createwallets

NODE_ID=4000 ./main addwallet
wallet address:  1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG

NODE_ID=4000 ./main createblockchain -address 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG

cp blockchain_4000.db blockchain_genesis.db

# 4001 钱包节点
NODE_ID=4001 ./main createwallets

NODE_ID=4001 ./main addwallet
wallet address:  12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx

NODE_ID=4001 ./main addwallet
wallet address:  19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6

NODE_ID=4001 ./main addwallet
wallet address:  1MqDHaopLxZrD4N1gLR8VaW3g7YgifbKTM

# 4000 shell 向4001钱包地址发送一些币
NODE_ID=4000 ./main send -from 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG -to 12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx -amount 3 -mine

NODE_ID=4000 ./main send -from 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG -to 19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6 -amount 1 -mine

NODE_ID=4000 ./main getbalance -address 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG
NODE_ID=4000 ./main getbalance -address 12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx
NODE_ID=4000 ./main getbalance -address 19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6

NODE_ID=4000 ./main startnode


# 4001 钱包节点
cp blockchain_genesis.db blockchain_4001.db

同步区块
NODE_ID=4001 ./main startnode

暂停所有节点，检查余额
NODE_ID=4000 ./main getbalance -address 12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx
NODE_ID=4001 ./main getbalance -address 12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx

NODE_ID=4000 ./main getbalance -address 19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6
NODE_ID=4001 ./main getbalance -address 19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6

NODE_ID=4000 ./main getbalance -address 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG
NODE_ID=4001 ./main getbalance -address 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG

# 4002 矿工节点
cp blockchain_genesis.db blockchain_4002.db

NODE_ID=4002 ./main createwallets

NODE_ID=4002 ./main addwallet
wallet address:  1N6VVzTNjYWFMdTpUiSHdcBN8vE6pBdr7J

启动节点
NODE_ID=4000 ./main startnode
NODE_ID=4002 ./main startnode -miner 1N6VVzTNjYWFMdTpUiSHdcBN8vE6pBdr7J

# 4001 钱包节点
NODE_ID=4001 ./main send -from 12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx -to 1MqDHaopLxZrD4N1gLR8VaW3g7YgifbKTM -amount 1

NODE_ID=4001 ./main send -from 19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6 -to 1MqDHaopLxZrD4N1gLR8VaW3g7YgifbKTM -amount 1

执行过程:
tip: 先同步三个节点的数据
1. sendTx to 4000
2. 4000 sendInv to 4002
3. 4002 handleInv, then sendGetData to 4000
4. 4000 handleGetData, get tx := mempool[txID] from mempool, then sendTx to 4002
5. 4002 handleTx, mempool > 2, miner

观察4000中心节点和4002矿工节点

NODE_ID=4001 ./main startnode 同步数据

暂停所有节点，检查余额
NODE_ID=4000 ./main getbalance -address 1MqDHaopLxZrD4N1gLR8VaW3g7YgifbKTM
NODE_ID=4001 ./main getbalance -address 1MqDHaopLxZrD4N1gLR8VaW3g7YgifbKTM
NODE_ID=4002 ./main getbalance -address 1MqDHaopLxZrD4N1gLR8VaW3g7YgifbKTM

NODE_ID=4000 ./main getbalance -address 12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx
NODE_ID=4001 ./main getbalance -address 12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx
NODE_ID=4002 ./main getbalance -address 12XSdrHBWCf4Nbvr8QopX7sLmczSd7ysJx

NODE_ID=4000 ./main getbalance -address 19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6
NODE_ID=4001 ./main getbalance -address 19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6
NODE_ID=4002 ./main getbalance -address 19iqa2dvZTLgwSfV1JRcwqcS4ADyZtHcd6

NODE_ID=4000 ./main getbalance -address 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG
NODE_ID=4001 ./main getbalance -address 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG
NODE_ID=4002 ./main getbalance -address 1QCCDfsssQbpYzXwQg5gDUnfLah5EGyEuG

NODE_ID=4000 ./main getbalance -address 1N6VVzTNjYWFMdTpUiSHdcBN8vE6pBdr7J
NODE_ID=4001 ./main getbalance -address 1N6VVzTNjYWFMdTpUiSHdcBN8vE6pBdr7J
NODE_ID=4002 ./main getbalance -address 1N6VVzTNjYWFMdTpUiSHdcBN8vE6pBdr7J
