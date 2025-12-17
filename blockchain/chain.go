package blockchain

import (
	"errors"
	"sync"
)

type blockChain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

// Single Pattern 으로 만들기
var bc *blockChain
var once sync.Once // 몇개의 채널이 있던 한번만 실행되도록 하기

func (bc *blockChain) AddBlock(data string) {
	block := createBlock(data, bc.NewestHash, bc.Height)
	bc.NewestHash = block.Hash
	bc.Height = block.Height
}

func GetBlockChain() *blockChain {
	if bc == nil {
		once.Do(func() {
			bc = &blockChain{"", 0}
			bc.AddBlock("Genesis Block")
		})
	}
	return bc
}

var ErrNotFound = errors.New("block not found")
