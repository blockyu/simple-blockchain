package core

import (
	"fmt"
	"simple-blockchain/consensus"
	"simple-blockchain/types"
	"time"
)

type Blockchain struct {
	blocks []*types.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) SearchBlock(from, to int) {
	for i, block := range bc.blocks {
		if i < from || i > to {
			continue
		}
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash. : %x\n", block.Hash)
		fmt.Printf("Timestamp. : %d\n\n", block.Timestamp)
	}
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*types.Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *types.Block {
	return NewBlock("Genesis Block generate !! ", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *types.Block {
	block := &types.Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := consensus.NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}
