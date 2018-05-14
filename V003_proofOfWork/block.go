package main

import (
	"time"
)

// Block 区块结构
type Block struct {
	Timestamp     int64  // 时间戳
	Data          []byte // 数据
	PrevBlockHash []byte // 上一个区块的Hash
	Hash          []byte // 区块Hash
	Nonce         int
}

// NewGenesisBlock 创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlock creates and returns Block
func NewBlock(data string, PrevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: PrevBlockHash,
		Hash:          []byte{}}

	// block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
