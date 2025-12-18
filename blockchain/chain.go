package blockchain

import (
	"fmt"
	"sync"

	"github.com/hellotect2022go/nomadcoin/db"
	"github.com/hellotect2022go/nomadcoin/utils"
)

type blockChain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

func (bc *blockChain) persist() {
	db.SaveBlockchain(utils.ToBytes(bc))
}

func (bc *blockChain) restore(data []byte) {
	utils.FromBytes(data, bc)
}

// Single Pattern 으로 만들기
var bc *blockChain
var once sync.Once // 몇개의 채널이 있던 한번만 실행되도록 하기

func (bc *blockChain) AddBlock(data string) {
	block := createBlock(data, bc.NewestHash, bc.Height+1)
	bc.NewestHash = block.Hash
	bc.Height = block.Height
	bc.persist()
}

func (bc *blockChain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := bc.NewestHash

	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func GetBlockChain() *blockChain {
	if bc == nil {
		once.Do(func() {
			bc = &blockChain{"", 0}
			//search for checkpoint on the db
			// restore bc from byte
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				bc.AddBlock("Genesis Block")
			} else {
				fmt.Println("Restoring....")
				bc.restore(checkpoint)
			}
		})
	}
	fmt.Printf("NewestHash: %s\nHeight: %d\n\n", bc.NewestHash, bc.Height)
	return bc
}
