package main

import (
	"crypto/sha256"
	"log"
	"strconv"
	"time"
)

// Block 区块结构
type Block struct {
	Timestamp     int64  // 时间戳
	Data          []byte // 数据
	PrevBlockHash []byte // 上一个区块的Hash
	Hash          []byte // 区块Hash
}

// NewGenesisBlock 创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte("0"))
}

// NewBlock creates and returns Block
func NewBlock(data string, PrevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: PrevBlockHash,
		Hash:          []byte("")}
	block.SetHash()
	return block
}

// SetHash 计算区块的Hash
func (b *Block) SetHash() {

	strTimeStamp := strconv.FormatInt(b.Timestamp, 10)
	timestamp := []byte(strTimeStamp)
	var data []byte
	data = append(data, b.PrevBlockHash...)
	data = append(data, b.Data...)
	data = append(data, timestamp...)

	h := sha256.New()
	_, err := h.Write(data)
	if err != nil {
		log.Panic(err)
	}

	b.Hash = h.Sum(nil)
}
