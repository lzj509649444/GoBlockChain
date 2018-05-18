package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const nodeVersion = 1

var knownNodes = []string{"localhost:4000"}

// verack The verack message is sent in reply to version. This message consists of only a message header with the command string "verack"
type verack struct {
}

// addr: Provide information on known nodes of the network
type addr struct {
	AddrList []string
}

// When a node creates an outgoing connection, it will immediately advertise its version. The remote node will respond with its version. No further communication is possible until both peers have exchanged their version
// version 用于找到一个更长的区块链。当一个节点接收到 version 消息，它会检查本节点的区块链是否比 BestHeight 的值更大。如果不是，节点就会请求并下载缺失的块。
type verzion struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

func sendVrack(addr string) {
	payload := gobEncode(verack{})

	request := append(commandToBytes("verack"), payload...)

	sendData(addr, request)
}

func sendAddr(address string) {
	nodes := addr{knownNodes}
	nodes.AddrList = append(nodes.AddrList, nodeAddress)
	payload := gobEncode(nodes)
	request := append(commandToBytes("addr"), payload...)

	sendData(address, request)
}

func handleAddr(request []byte) {
	var buff bytes.Buffer
	var payload addr

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	knownNodes = append(knownNodes, payload.AddrList...)
	fmt.Printf("There are %d known nodes now!\n", len(knownNodes))
}

// GetBestHeight returns the height of the latest block
func (blockchain *Blockchain) GetBestHeight() int {
	var lastBlock Block

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash := b.Get([]byte("l"))
		blockData := b.Get(lastHash)
		lastBlock = *DeserializeBlock(blockData)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return lastBlock.Height
}

func sendVersion(addr string, blockchain *Blockchain) {
	bestHeight := blockchain.GetBestHeight()
	payload := gobEncode(verzion{nodeVersion, bestHeight, nodeAddress})

	request := append(commandToBytes("version"), payload...)

	sendData(addr, request)
}

func handleVersion(request []byte, blockchain *Blockchain) {
	var buff bytes.Buffer
	var payload verzion

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	myBestHeight := blockchain.GetBestHeight()
	foreignerBestHeight := payload.BestHeight

	if myBestHeight < foreignerBestHeight {
		sendGetBlocks(payload.AddrFrom)
	} else if myBestHeight > foreignerBestHeight {
		sendVersion(payload.AddrFrom, blockchain)
	}

	if !nodeIsKnown(payload.AddrFrom) {
		knownNodes = append(knownNodes, payload.AddrFrom)
	}
	fmt.Printf("knownNodes %s\n", knownNodes)

	// sendVrack(payload.AddrFrom)
	// sendAddr(payload.AddrFrom)
	// knownNodes = append(knownNodes, payload.AddrFrom)
}

func nodeIsKnown(addr string) bool {
	for _, node := range knownNodes {
		if node == addr {
			return true
		}
	}

	return false
}
