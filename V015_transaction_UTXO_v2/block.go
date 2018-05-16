package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Block 区块结构
type Block struct {
	Timestamp     int64 // 时间戳
	Transactions  []*Transaction
	PrevBlockHash []byte // 上一个区块的Hash
	Hash          []byte // 区块Hash
	Nonce         int
}

func (block *Block) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("{Timestamp: %d", block.Timestamp))
	lines = append(lines, fmt.Sprintf("Transactions: %s", block.Transactions))
	lines = append(lines, fmt.Sprintf("PrevBlockHash: %x", block.PrevBlockHash))
	lines = append(lines, fmt.Sprintf("Hash: %x", block.Hash))
	lines = append(lines, fmt.Sprintf("Nonce: %x}", block.Nonce))

	pow := NewProofOfWork(block)
	lines = append(lines, fmt.Sprintf("PoW: %s", strconv.FormatBool(pow.Validate())))

	return strings.Join(lines, "\n")
}

// HashTransactions returns a hash of the transactions in the block
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// NewGenesisBlock 创世区块
func NewGenesisBlock(coinbase *Transaction) *Block {
	transactions := []*Transaction{coinbase}
	return NewBlock(transactions, []byte{})
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, PrevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: PrevBlockHash,
		Hash:          []byte{}}

	// block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	fmt.Printf("Block: %s\n", pow.block)

	return block
}
