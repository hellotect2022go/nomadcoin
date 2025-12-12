package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockChain struct {
	blocks []block
}

func (bc *blockChain) getLastHash() string {
	if len(bc.blocks) > 0 {
		return bc.blocks[len(bc.blocks)-1].hash
	}
	return ""
}

func (bc *blockChain) addBlock(data string) {
	newBlock := block{data, "", bc.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *blockChain) listBlocks() {
	for index, block := range bc.blocks {
		fmt.Printf("[%d]================\n", index)
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("PrevHash: %s\n", block.prevHash)
	}
}

func main() {

	chain := blockChain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}
