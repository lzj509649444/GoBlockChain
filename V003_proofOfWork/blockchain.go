package main

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain genesisBlock chain
func NewBlockchain() *Blockchain {
	blockchain := Blockchain{
		blocks: []*Block{NewGenesisBlock()}}
	return &blockchain
}

// AddBlock add newBlock
func (blockChain *Blockchain) AddBlock(data string) {
	prevIndex := len(blockChain.blocks) - 1
	prevBlock := blockChain.blocks[prevIndex]
	newBlock := NewBlock(data, prevBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, newBlock)
}
