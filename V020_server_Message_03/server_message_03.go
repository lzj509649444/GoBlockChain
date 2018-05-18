package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

var miningAddress string
var mempool = make(map[string]Transaction)

type tx struct {
	AddFrom     string
	Transaction []byte
}

func sendTx(addr string, transaction *Transaction) {
	data := tx{nodeAddress, transaction.Serialize()}
	payload := gobEncode(data)
	request := append(commandToBytes("tx"), payload...)

	sendData(addr, request)
}

func handleTx(request []byte, blockchain *Blockchain) {
	var buff bytes.Buffer
	var payload tx

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	txData := payload.Transaction
	tx := DeserializeTransaction(txData)
	mempool[hex.EncodeToString(tx.ID)] = tx

	// 检查当前节点是否是中心节点。在我们的实现中，中心节点并不会挖矿。它只会将新的交易推送给网络中的其他节点。
	if nodeAddress == knownNodes[0] {
		for _, node := range knownNodes {
			if node != nodeAddress && node != payload.AddFrom {
				sendInv(node, "tx", [][]byte{tx.ID})
			}
		}
	} else {
		// 矿工节点“专属”
		// miningAddress 只会在矿工节点上设置。如果当前节点（矿工）的内存池中有两笔或更多的交易，开始挖矿
		if len(mempool) >= 2 && len(miningAddress) > 0 {
		MineTransactions:
			var txs []*Transaction

			for id := range mempool {
				tx := mempool[id]
				if blockchain.VerifyTransaction(&tx) {
					txs = append(txs, &tx)
				}
			}

			if len(txs) == 0 {
				fmt.Println("All transactions are invalid! Waiting for new ones...")
				return
			}

			// 验证后的交易被放到一个块里，同时还有附带奖励的 coinbase 交易
			cbTx := NewCoinbaseTX(miningAddress, "")
			txs = append(txs, cbTx)

			newBlock := blockchain.MineBlock(txs)
			utxoSet := UTXOSet{blockchain}
			// TODO: 提醒，应该使用 UTXOSet.Update 而不是 UTXOSet.Reindex.
			utxoSet.Reindex()

			fmt.Println("New block is mined!")

			for _, tx := range txs {
				txID := hex.EncodeToString(tx.ID)
				delete(mempool, txID)
			}

			// 当一笔交易被挖出来以后，就会被从内存池中移除。当前节点所连接到的所有其他节点，接收带有新块哈希的 inv 消息。在处理完消息后，它们可以对块进行请求。
			for _, node := range knownNodes {
				if node != nodeAddress {
					fmt.Printf("%s handleTx send block to %s\n", nodeAddress, node)
					sendInv(node, "block", [][]byte{newBlock.Hash})
				}
			}

			if len(mempool) > 0 {
				goto MineTransactions
			}
		}
	}
}
