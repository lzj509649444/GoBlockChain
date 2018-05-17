package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var blocksInTransit = [][]byte{}

// getblocks 意为 “给我看一下你有什么区块”（在比特币中，这会更加复杂）。注意，它并没有说“把你全部的区块给我”，而是请求了一个块哈希的列表。这是为了减轻网络负载，因为区块可以从不同的节点下载，并且我们不想从一个单一节点下载数十 GB 的数据
type getblocks struct {
	AddrFrom string
}

// 比特币使用 inv 来向其他节点展示当前节点有什么块和交易。再次提醒，它没有包含完整的区块链和交易，仅仅是哈希而已。Type 字段表明了这是块还是交易。
// Allows a node to advertise its knowledge of one or more objects. It can be received unsolicited, or in reply to getblocks
type inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

// getdata 用于某个块或交易的请求，它可以仅包含一个块或交易的 ID
// getdata is used in response to inv, to retrieve the content of a specific object, and is usually sent after receiving an inv packet, after filtering known elements. It can be used to retrieve transactions, but only if they are in the memory pool or relay set - arbitrary access to transactions in the chain is not allowed to avoid having clients start to depend on nodes having full transaction indexes (which modern nodes do not).
type getdata struct {
	AddrFrom string
	Type     string
	ID       []byte
}

// The block message is sent in response to a getdata message which requests transaction information from a block hash
type block struct {
	AddrFrom string
	Block    []byte
}

// GetBlockHashes returns a list of hashes of all the blocks in the chain
func (blockchain *Blockchain) GetBlockHashes() [][]byte {
	var hashes [][]byte
	iter := blockchain.Iterator()

	for {
		block := iter.Next()

		hashes = append(hashes, block.Hash)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return hashes
}

func sendGetBlocks(address string) {
	payload := gobEncode(getblocks{nodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	sendData(address, request)
}

func handleGetBlocks(request []byte, blockchain *Blockchain) {
	var buff bytes.Buffer
	var payload getblocks

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blocks := blockchain.GetBlockHashes()
	sendInv(payload.AddrFrom, "blocks", blocks)
}

func sendInv(address, kind string, items [][]byte) {
	// 库存
	inventory := inv{nodeAddress, kind, items}
	payload := gobEncode(inventory)
	request := append(commandToBytes("inv"), payload...)

	sendData(address, request)
}

// 如果收到块哈希，我们想要将它们保存在 blocksInTransit 变量来跟踪已下载的块。这能够让我们从不同的节点下载块。在将块置于传送状态时，我们给 inv 消息的发送者发送 getdata 命令并更新 blocksInTransit。在一个真实的 P2P 网络中，我们会想要从不同节点来传送块。
// 在我们的实现中，我们永远也不会发送有多重哈希的 inv。这就是为什么当 payload.Type == "tx" 时，只会拿到第一个哈希。然后我们检查是否在内存池中已经有了这个哈希，如果没有，发送 getdata 消息。
func handleInv(request []byte, blockchain *Blockchain) {
	var buff bytes.Buffer
	var payload inv

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Recevied inventory with %d %s\n", len(payload.Items), payload.Type)

	if payload.Type == "blocks" {
		blocksInTransit = payload.Items

		blockHash := payload.Items[0]
		sendGetData(payload.AddrFrom, "block", blockHash)

		newInTransit := [][]byte{}
		for _, b := range blocksInTransit {
			if bytes.Compare(b, blockHash) != 0 {
				newInTransit = append(newInTransit, b)
			}
		}
		blocksInTransit = newInTransit
	}
}

func sendGetData(address, kind string, id []byte) {
	payload := gobEncode(getdata{nodeAddress, kind, id})
	request := append(commandToBytes("getdata"), payload...)

	sendData(address, request)
}

func handleGetData(request []byte, blockchain *Blockchain) {
	var buff bytes.Buffer
	var payload getdata

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	if payload.Type == "block" {
		block, err := blockchain.GetBlock([]byte(payload.ID))
		if err != nil {
			return
		}
		sendBlock(payload.AddrFrom, &block)
	}
}

// GetBlock finds a block by its hash and returns it
func (blockchain *Blockchain) GetBlock(blockHash []byte) (Block, error) {
	var block Block

	err := blockchain.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		blockData := b.Get(blockHash)

		if blockData == nil {
			return errors.New("block is not found")
		}

		block = *DeserializeBlock(blockData)

		return nil
	})
	if err != nil {
		return block, err
	}

	return block, nil
}

func sendBlock(address string, b *Block) {
	data := block{nodeAddress, b.Serialize()}
	payload := gobEncode(data)
	request := append(commandToBytes("block"), payload...)

	sendData(address, request)
}

func handleBlock(request []byte, blockchain *Blockchain) {
	var buff bytes.Buffer
	var payload block

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockData := payload.Block
	block := DeserializeBlock(blockData)

	fmt.Println("Recevied a new block!")
	blockchain.AddBlock(block)

	fmt.Println(blocksInTransit)
	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		sendGetData(payload.AddrFrom, "block", blockHash)

		blocksInTransit = blocksInTransit[1:]
	}
}

// AddBlock saves the block into the blockchain
func (blockchain *Blockchain) AddBlock(block *Block) {
	err := blockchain.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		blockInDb := b.Get(block.Hash)

		if blockInDb != nil {
			return nil
		}

		blockData := block.Serialize()
		err := b.Put(block.Hash, blockData)
		if err != nil {
			log.Panic(err)
		}

		lastHash := b.Get([]byte("l"))
		lastBlockData := b.Get(lastHash)
		lastBlock := DeserializeBlock(lastBlockData)

		if block.Height > lastBlock.Height {
			err = b.Put([]byte("l"), block.Hash)
			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
