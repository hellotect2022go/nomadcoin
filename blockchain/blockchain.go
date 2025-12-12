package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockChain struct {
	blocks []*block
}

// Single Pattern 으로 만들기
var bc *blockChain
var once sync.Once // 몇개의 채널이 있던 한번만 실행되도록 하기

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash) //16진수 : base16 문자열로 전환
}

func getLastHash() string {
	totalBlocks := len(GetBlockChain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockChain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockChain() *blockChain {
	if bc == nil {
		once.Do(func() {
			bc = &blockChain{}
			bc.AddBlock("Genesis Block")
		})
	}
	return bc
}

//

func AllBlocks() []*block {
	return GetBlockChain().blocks
}
